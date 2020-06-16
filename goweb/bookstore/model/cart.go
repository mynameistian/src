package model

//Cart 购物车结构体
type Cart struct {
	CartID      string      //购物车的id
	CartItems   []*CartItem //购物车项
	TotalCount  int64       //购物车中图书总数量，通过计算得到
	TotalAmount float64     //购物车中的图书的总金额，通过计算得到
	UserID      int
}

//GetTotalCount 计算总数量
func (cart *Cart) GetTotalCount() int64 {
	var totalCount int64
	for _, v := range cart.CartItems {
		totalCount = totalCount + v.Count
	}
	return totalCount
}

//GetTotalAmount 计算总金额
func (cart *Cart) GetTotalAmount() float64 {
	var totalAmount float64
	for _, v := range cart.CartItems {
		totalAmount = totalAmount + v.GetAmount()
	}
	return totalAmount
}
