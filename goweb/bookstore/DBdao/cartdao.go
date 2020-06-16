package DBdao

import (
	"goweb/bookstore/model"
	"goweb/bookstore/utils"
)

//AddCart 向购物车表中插入购物车
func AddCart(cart *model.Cart) error {

	sqlStr := "insert into carts(id,total_count,total_amount,user_id)values($1,$2,$3,$4);"

	_, err := utils.Db.Exec(sqlStr, cart.CartID, cart.GetTotalCount(), cart.GetTotalAmount(), cart.UserID)
	if err != nil {
		return err
	}

	//获取购物车中的所有购物项
	cartItms := cart.CartItems
	//遍历得到每一个购物项插入购物项表中
	for _, cartItem := range cartItms {
		AddCartItem(cartItem)
	}

	return nil
}

//GetCartByUserId
func GetCartByUserId(userID int) (*model.Cart, error) {
	sqlStr := "select id ,total_count,total_amount,user_id from carts where user_id = $1;"
	Row := utils.Db.QueryRow(sqlStr, userID)
	cart := &model.Cart{}
	err := Row.Scan(&cart.CartID, &cart.TotalCount, &cart.TotalAmount, &cart.UserID)
	if err != nil {
		return nil, err
	}
	cart.CartItems, _ = GetCatrtItemByCartID(cart.CartID)
	return cart, nil
}

//UpdateCart 更新购物车中个的总数量及总金额
func UpdateCart(cart *model.Cart) error {
	sqlStr := "update carts set total_count = $1 , total_amount = $2 where id = $3;"

	_, err := utils.Db.Exec(sqlStr, cart.GetTotalCount(), cart.GetTotalAmount(), cart.CartID)
	if err != nil {
		return err
	}
	return nil
}
