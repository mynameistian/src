// Copyright 2019 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package lsp

import (
	"bytes"
	"context"
	"fmt"

	"golang.org/x/tools/internal/jsonrpc2"
	"golang.org/x/tools/internal/lsp/protocol"
	"golang.org/x/tools/internal/lsp/source"
	"golang.org/x/tools/internal/span"
	errors "golang.org/x/xerrors"
)

func (s *Server) didOpen(ctx context.Context, params *protocol.DidOpenTextDocumentParams) error {
	uri := params.TextDocument.URI.SpanURI()
	if !uri.IsFile() {
		return nil
	}

	_, err := s.didModifyFiles(ctx, []source.FileModification{
		{
			URI:        uri,
			Action:     source.Open,
			Version:    params.TextDocument.Version,
			Text:       []byte(params.TextDocument.Text),
			LanguageID: params.TextDocument.LanguageID,
		},
	})
	return err
}

func (s *Server) didChange(ctx context.Context, params *protocol.DidChangeTextDocumentParams) error {
	uri := params.TextDocument.URI.SpanURI()
	if !uri.IsFile() {
		return nil
	}

	text, err := s.changedText(ctx, uri, params.ContentChanges)
	if err != nil {
		return err
	}
	c := source.FileModification{
		URI:     uri,
		Action:  source.Change,
		Version: params.TextDocument.Version,
		Text:    text,
	}
	snapshots, err := s.didModifyFiles(ctx, []source.FileModification{c})
	if err != nil {
		return err
	}
	snapshot := snapshots[uri]
	if snapshot == nil {
		return errors.Errorf("no snapshot for %s", uri)
	}
	// Ideally, we should be able to specify that a generated file should be opened as read-only.
	// Tell the user that they should not be editing a generated file.
	if s.wasFirstChange(uri) && source.IsGenerated(ctx, snapshot, uri) {
		if err := s.client.ShowMessage(ctx, &protocol.ShowMessageParams{
			Message: fmt.Sprintf("Do not edit this file! %s is a generated file.", uri.Filename()),
			Type:    protocol.Warning,
		}); err != nil {
			return err
		}
	}
	return nil
}

func (s *Server) didChangeWatchedFiles(ctx context.Context, params *protocol.DidChangeWatchedFilesParams) error {
	var modifications []source.FileModification
	deletions := make(map[span.URI]struct{})
	for _, change := range params.Changes {
		uri := change.URI.SpanURI()
		if !uri.IsFile() {
			continue
		}
		action := changeTypeToFileAction(change.Type)
		modifications = append(modifications, source.FileModification{
			URI:    uri,
			Action: action,
			OnDisk: true,
		})
		// Keep track of deleted files so that we can clear their diagnostics.
		// A file might be re-created after deletion, so only mark files that
		// have truly been deleted.
		switch action {
		case source.Delete:
			deletions[uri] = struct{}{}
		case source.Close:
		default:
			delete(deletions, uri)
		}
	}
	snapshots, err := s.didModifyFiles(ctx, modifications)
	if err != nil {
		return err
	}
	// Clear the diagnostics for any deleted files.
	for uri := range deletions {
		if snapshot := snapshots[uri]; snapshot == nil || snapshot.IsOpen(uri) {
			continue
		}
		if err := s.client.PublishDiagnostics(ctx, &protocol.PublishDiagnosticsParams{
			URI:         protocol.URIFromSpanURI(uri),
			Diagnostics: []protocol.Diagnostic{},
			Version:     0,
		}); err != nil {
			return err
		}
	}
	return nil
}

func (s *Server) didSave(ctx context.Context, params *protocol.DidSaveTextDocumentParams) error {
	uri := params.TextDocument.URI.SpanURI()
	if !uri.IsFile() {
		return nil
	}
	c := source.FileModification{
		URI:     uri,
		Action:  source.Save,
		Version: params.TextDocument.Version,
	}
	if params.Text != nil {
		c.Text = []byte(*params.Text)
	}
	_, err := s.didModifyFiles(ctx, []source.FileModification{c})
	return err
}

func (s *Server) didClose(ctx context.Context, params *protocol.DidCloseTextDocumentParams) error {
	uri := params.TextDocument.URI.SpanURI()
	if !uri.IsFile() {
		return nil
	}
	snapshots, err := s.didModifyFiles(ctx, []source.FileModification{
		{
			URI:     uri,
			Action:  source.Close,
			Version: -1,
			Text:    nil,
		},
	})
	if err != nil {
		return err
	}
	snapshot := snapshots[uri]
	if snapshot == nil {
		return errors.Errorf("no snapshot for %s", uri)
	}
	fh, err := snapshot.GetFile(uri)
	if err != nil {
		return err
	}
	// If a file has been closed and is not on disk, clear its diagnostics.
	if _, _, err := fh.Read(ctx); err != nil {
		return s.client.PublishDiagnostics(ctx, &protocol.PublishDiagnosticsParams{
			URI:         protocol.URIFromSpanURI(uri),
			Diagnostics: []protocol.Diagnostic{},
			Version:     0,
		})
	}
	return nil
}

func (s *Server) didModifyFiles(ctx context.Context, modifications []source.FileModification) (map[span.URI]source.Snapshot, error) {
	snapshots, err := s.session.DidModifyFiles(ctx, modifications)
	if err != nil {
		return nil, err
	}
	snapshotByURI := make(map[span.URI]source.Snapshot)
	for _, c := range modifications {
		snapshotByURI[c.URI] = nil
	}
	// Avoid diagnosing the same snapshot twice.
	snapshotSet := make(map[source.Snapshot][]span.URI)
	for uri := range snapshotByURI {
		view, err := s.session.ViewOf(uri)
		if err != nil {
			return nil, err
		}
		var snapshot source.Snapshot
		for _, s := range snapshots {
			if s.View() == view {
				if snapshot != nil {
					return nil, errors.Errorf("duplicate snapshots for the same view")
				}
				snapshot = s
			}
		}
		// If the file isn't in any known views (for example, if it's in a dependency),
		// we may not have a snapshot to map it to. As a result, we won't try to
		// diagnose it. TODO(rstambler): Figure out how to handle this better.
		if snapshot == nil {
			continue
		}
		snapshotByURI[uri] = snapshot
		snapshotSet[snapshot] = append(snapshotSet[snapshot], uri)
	}
	for snapshot, uris := range snapshotSet {
		// If a modification comes in for the view's go.mod file and the view
		// was never properly initialized, or the view does not have
		// a go.mod file, try to recreate the associated view.
		if modfile, _ := snapshot.View().ModFiles(); modfile == "" {
			for _, uri := range uris {
				// Don't rebuild the view until the go.mod is on disk.
				if !snapshot.IsSaved(uri) {
					continue
				}
				fh, err := snapshot.GetFile(uri)
				if err != nil {
					return nil, err
				}
				switch fh.Identity().Kind {
				case source.Mod:
					newSnapshot, err := snapshot.View().Rebuild(ctx)
					if err != nil {
						return nil, err
					}
					// Update the snapshot to the rebuilt one.
					snapshot = newSnapshot
					snapshotByURI[uri] = newSnapshot
				}
			}
		}
		go s.diagnoseSnapshot(snapshot)
	}
	return snapshotByURI, nil
}

func (s *Server) wasFirstChange(uri span.URI) bool {
	if s.changedFiles == nil {
		s.changedFiles = make(map[span.URI]struct{})
	}
	_, ok := s.changedFiles[uri]
	return ok
}

func (s *Server) changedText(ctx context.Context, uri span.URI, changes []protocol.TextDocumentContentChangeEvent) ([]byte, error) {
	if len(changes) == 0 {
		return nil, jsonrpc2.NewErrorf(jsonrpc2.CodeInternalError, "no content changes provided")
	}

	// Check if the client sent the full content of the file.
	// We accept a full content change even if the server expected incremental changes.
	if len(changes) == 1 && changes[0].Range == nil && changes[0].RangeLength == 0 {
		return []byte(changes[0].Text), nil
	}
	return s.applyIncrementalChanges(ctx, uri, changes)
}

func (s *Server) applyIncrementalChanges(ctx context.Context, uri span.URI, changes []protocol.TextDocumentContentChangeEvent) ([]byte, error) {
	content, _, err := s.session.GetFile(uri).Read(ctx)
	if err != nil {
		return nil, jsonrpc2.NewErrorf(jsonrpc2.CodeInternalError, "file not found (%v)", err)
	}
	for _, change := range changes {
		// Make sure to update column mapper along with the content.
		converter := span.NewContentConverter(uri.Filename(), content)
		m := &protocol.ColumnMapper{
			URI:       uri,
			Converter: converter,
			Content:   content,
		}
		if change.Range == nil {
			return nil, jsonrpc2.NewErrorf(jsonrpc2.CodeInternalError, "unexpected nil range for change")
		}
		spn, err := m.RangeSpan(*change.Range)
		if err != nil {
			return nil, err
		}
		if !spn.HasOffset() {
			return nil, jsonrpc2.NewErrorf(jsonrpc2.CodeInternalError, "invalid range for content change")
		}
		start, end := spn.Start().Offset(), spn.End().Offset()
		if end < start {
			return nil, jsonrpc2.NewErrorf(jsonrpc2.CodeInternalError, "invalid range for content change")
		}
		var buf bytes.Buffer
		buf.Write(content[:start])
		buf.WriteString(change.Text)
		buf.Write(content[end:])
		content = buf.Bytes()
	}
	return content, nil
}

func changeTypeToFileAction(ct protocol.FileChangeType) source.FileAction {
	switch ct {
	case protocol.Changed:
		return source.Change
	case protocol.Created:
		return source.Create
	case protocol.Deleted:
		return source.Delete
	}
	return source.UnknownFileAction
}
