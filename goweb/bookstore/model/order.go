package model

type Order struct {
	ID          string
	CreateTime  string
	TotalCount  int64
	TotalAmount float64
	State       int64
	Uer_id      int64
}
