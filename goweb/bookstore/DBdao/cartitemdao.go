package DBdao

import (
	"goweb/bookstore/model"
	"goweb/bookstore/utils"
)

func AddCartItem(cartItem *model.CartItem) error {
	sqlStr := "insert into cart_itmes(count,amount,book_id,cart_id)values($1,$2,$3,$4);"

	_, err := utils.Db.Exec(sqlStr, cartItem.Count, cartItem.GetAmount(), cartItem.Book.ID, cartItem.CartID)
	if err != nil {
		return err
	}
	return nil
}

//GetCartItemByBookID 根据图书的id获取对应的购物项
func GetCartItemByBookID(bookID string) (*model.CartItem, error) {
	sqlStr := "select id ,count , amount ,cart_id from cart_itmes where book_id = $1;"

	row := utils.Db.QueryRow(sqlStr, bookID)

	cartItem := &model.CartItem{}
	err := row.Scan(&cartItem.CartItemID, &cartItem.Count, &cartItem.Amount, &cartItem.CartID)
	if err != nil {
		return nil, err
	}

	return cartItem, nil
}

//GetCatrtItemByCartID 根据购物车的id获取购物车中的所有购物项
func GetCatrtItemByCartID(cartID string) ([]*model.CartItem, error) {
	sqlStr := "select id ,count , amount ,cart_id ,book_id from cart_itmes where cart_id = $1;"

	rows, err := utils.Db.Query(sqlStr, cartID)
	if err != nil {
		return nil, err
	}

	var cartItems []*model.CartItem

	for rows.Next() {

		var bookId int
		cartItem := &model.CartItem{}
		err := rows.Scan(&cartItem.CartItemID, &cartItem.Count, &cartItem.Amount, &cartItem.CartID, &bookId)
		if err != nil {
			return nil, err
		}
		book, err := GetBookByIdID(bookId)
		if err != nil {
			return nil, err
		}
		cartItem.Book = book
		cartItems = append(cartItems, cartItem)
	}

	return cartItems, nil
}

//GetCartItemByBookIDAndCardID 根据图书的id和购物车id 获取对应的购物项
func GetCartItemByBookIDAndCardID(bookID string, cartID string) (*model.CartItem, error) {
	sqlStr := "select id ,count , amount ,cart_id from cart_itmes where book_id = $1 and cart_id = $2;"

	row := utils.Db.QueryRow(sqlStr, bookID, cartID)

	cartItem := &model.CartItem{}
	err := row.Scan(&cartItem.CartItemID, &cartItem.Count, &cartItem.Amount, &cartItem.CartID)
	if err != nil {
		return nil, err
	}

	return cartItem, nil
}
