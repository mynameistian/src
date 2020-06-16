package model

//CartItem 购物车结构体
type CartItem struct {
	CartItemID int64   //购物项id
	Book       *Book   //购物项中的图书信息
	Count      int64   //购物项中的图书数量
	Amount     float64 //购物项中图书的金额小计，通过计算得到
	CartID     string  //当前购物项属于哪个购物车
}

//获取金额小计的方法
func (cartitem *CartItem) GetAmount() float64 {
	price := cartitem.Book.Price

	cartitem.Amount = float64(cartitem.Count) * price
	return cartitem.Amount

}
