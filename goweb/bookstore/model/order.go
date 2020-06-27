package model

type Order struct {
	ID          string
	CreateTime  string
	TotalCount  int64
	TotalAmount float64
	State       string
	Uer_id      int64
}
