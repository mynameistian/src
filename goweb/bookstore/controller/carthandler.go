package controller

import (
	_ "fmt"
	"goweb/bookstore/DBdao"
	"goweb/bookstore/model"
	"goweb/bookstore/utils"
	_ "html/template"
	"net/http"
	"strconv"
)

func AddBook2Cart(w http.ResponseWriter, r *http.Request) {
	bookID := r.FormValue("bookId")
	//fmt.Println(bookId)
	ibookId, _ := strconv.Atoi(bookID)
	book, _ := DBdao.GetBookByIdID(ibookId)

	_, session := DBdao.IsLogin(r)

	userID := session.UserID

	cart, _ := DBdao.GetCartByUserId(userID)
	if cart != nil {
		//当前用户已经有购物车
		cartItem, _ := DBdao.GetCartItemByBookIDAndCardID(bookID, cart.CartID)
		if cartItem != nil {

		} else {
			cartItem := &model.CartItem{
				Book:   book,
				Count:  1,
				CartID: cart.CartID,
			}
			cart.CartItems = append(cart.CartItems, cartItem)
			//添加购物项
			DBdao.AddCartItem(cartItem)
		}
		//更新购物车总数据量和总金额信息
		err := DBdao.UpdateCart(cart)
		if err != nil {

		}
	} else {
		//当前用户还没有购物车
		cartID := utils.CreateUUID()
		cart := &model.Cart{
			CartID: cartID,
			UserID: userID,
		}

		var cartItems []*model.CartItem
		cartItem := &model.CartItem{
			Book:   book,
			Count:  1,
			CartID: cartID,
		}
		cartItems = append(cartItems, cartItem)
		cart.CartItems = cartItems
		DBdao.AddCart(cart)
	}

}
