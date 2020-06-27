package model

type Session struct {
	SessionID string
	UserName  string
	UserID    int
	Cart      *Cart
	Order     *Order
	Orders    []*Order
}
