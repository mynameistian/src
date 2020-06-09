package main

import (
	"bufio"
	"fmt"
	"github.com/lxn/walk"
	. "github.com/lxn/walk/declarative"
	_ "io"
	"io/ioutil"
	_ "log"
	"os"
	"path/filepath"
	_ "strings"
)

// var le *walk.LineEdit
// var sport *walk.CheckBox

type MyMainWindow struct {
	*walk.MainWindow
	edit_fileName      *walk.TextEdit
	edit_fileName_flag bool
	edit_Key           *walk.TextEdit
	Label              *walk.LinkLabel
	fileName           string
}

func main() {

	mw := &MyMainWindow{}

	err := MainWindow{
		AssignTo: &mw.MainWindow,
		Title:    "文件选择框对话框",
		MinSize:  Size{Width: 300, Height: 200},
		MaxSize:  Size{Width: 600, Height: 200},
		Size:     Size{Width: 600, Height: 200},
		Layout:   VBox{},
		Children: []Widget{
			GroupBox{
				Layout: HBox{},
				Children: []Widget{
					LinkLabel{
						Text:    "文件路径:",
						MinSize: Size{Width: 60, Height: 20},
						MaxSize: Size{Width: 60, Height: 20},
					},
					TextEdit{
						Text:     "请选择文件",
						AssignTo: &mw.edit_fileName,
						MinSize:  Size{Width: 400, Height: 20},
						MaxSize:  Size{Width: 400, Height: 20},
					},
					PushButton{
						Text:      "打开文件",
						MinSize:   Size{Width: 60, Height: 40},
						MaxSize:   Size{Width: 60, Height: 40},
						OnClicked: mw.selectFile,
					},
				},
				MinSize: Size{Width: 550, Height: 60},
				MaxSize: Size{Width: 550, Height: 60},
			},
			GroupBox{
				Layout: HBox{},
				Children: []Widget{
					LinkLabel{
						Text:    "key:",
						MinSize: Size{Width: 60, Height: 20},
						MaxSize: Size{Width: 60, Height: 20},
					},
					TextEdit{
						Text:     "",
						AssignTo: &mw.edit_Key,
						MinSize:  Size{Width: 400, Height: 20},
						MaxSize:  Size{Width: 400, Height: 20},
					},
					PushButton{
						MinSize:   Size{Width: 60, Height: 60},
						MaxSize:   Size{Width: 60, Height: 60},
						Text:      "解密",
						OnClicked: mw.saveFile,
					},
				},
				MinSize: Size{Width: 550, Height: 60},
				MaxSize: Size{Width: 550, Height: 60},
			},
			LinkLabel{
				Text:     "执行结果",
				AssignTo: &mw.Label,
				MinSize:  Size{Width: 550, Height: 60},
				MaxSize:  Size{Width: 550, Height: 60},
			},
		},
	}.Create()

	if err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
	mw.Run()

}

func (mw *MyMainWindow) selectFile() {
	dlg := new(walk.FileDialog)
	dlg.Title = "选择文件"
	dlg.Filter = "所有文件(*.*)|*.*"

	mw.edit_fileName.SetText("")
	mw.edit_fileName_flag = false
	if ok, err := dlg.ShowOpen(mw); err != nil {
		mw.edit_fileName.SetText("Error: File Open\r\n")
		return
	} else if !ok {
		mw.edit_fileName.SetText("未选择文件\r\n")
		return
	}
	mw.edit_fileName_flag = true
	mw.fileName = dlg.FilePath
	s := fmt.Sprintf(" %s\r\n", dlg.FilePath)
	mw.edit_fileName.SetText(s)

}

//自己编写一个函数，接收两个文件路径 srcFileName dstFileName
func (mw *MyMainWindow) CopyFile(srcFileName string, dstFileName string) (written int, err error) {

	content, err := ioutil.ReadFile(srcFileName)
	if err != nil {
		fmt.Printf("read file err=%v", err)
	}

	//打开dstFileName
	dstFile, err := os.OpenFile(dstFileName, os.O_WRONLY|os.O_CREATE, 0666)
	if err != nil {
		fmt.Printf("open file err=%v\n", err)
		return
	}
	defer dstFile.Close()

	//通过dstFile, 获取到 Writer
	writer := bufio.NewWriter(dstFile)
	//输出内容
	key := mw.edit_Key.Text()
	encrypt_stream(content, len(content), key, 4096)
	writer.Write(content)
	writer.Flush()
	return
}

func (mw *MyMainWindow) saveFile() {

	if !mw.edit_fileName_flag {
		mw.Label.SetText("通知【没有选择文件】")
		return
	}
	if mw.edit_Key.TextLength() == 0 {
		mw.Label.SetText("通知【没有输入key】")
		return
	}
	//fileName := "F:\\桌面\\web下载\\123.xls"
	filePwd, fileName := filepath.Split(mw.fileName)
	outfileName := filePwd + "New" + fileName
	_, err := mw.CopyFile(mw.fileName, outfileName)
	if err != nil {
		fmt.Println("ReadBlck err ", err)
		mw.Label.SetText("执行失败")
		return
	}
	mw.Label.SetText("通知【解密完成，文件名称为：" + outfileName + "】")
}
