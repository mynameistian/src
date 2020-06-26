package controller

import (
	_ "fmt"
	"goweb/bookstore/DBdao"
	"goweb/bookstore/model"
	"goweb/bookstore/utils"
	"html/template"
	"net/http"
	"strconv"
)

func AddBook2Cart(w http.ResponseWriter, r *http.Request) {

	var userID int
	bLogin, session := DBdao.IsLogin(r)
	if bLogin {
		userID = session.UserID

		bookID := r.FormValue("bookId")
		//fmt.Println(bookId)
		ibookId, _ := strconv.Atoi(bookID)
		book, _ := DBdao.GetBookByIdID(ibookId)

		cart, _ := DBdao.GetCartByUserId(userID)
		if cart != nil {
			//当前用户已经有购物车
			//fmt.Println(bookID, " ", cart.CartID)
			cartItem, _ := DBdao.GetCartItemByBookIDAndCardID(bookID, cart.CartID)
			//fmt.Println(cartItem)
			if cartItem != nil {
				//当前这本书已有这本书
				cts := cart.CartItems
				for _, v := range cts {
					if v.Book.ID == cartItem.Book.ID {
						v.Count = v.Count + 1
						DBdao.UpdateBookCount(v)
					}
				}
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
			DBdao.UpdateCart(cart)
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
		w.Write([]byte("您刚刚将" + book.Title + "添加到了购物车！"))

	} else {
		w.Write([]byte("noLogin"))
	}
}

//GetCartInfo
func GetCartInfo(w http.ResponseWriter, r *http.Request) {
	bLogin, session := DBdao.IsLogin(r)
	if bLogin {
		userID := session.UserID
		cart, _ := DBdao.GetCartByUserId(userID)
		if cart != nil {
			t := template.Must(template.ParseFiles("views/pages/cart/cart.html"))
			cart.UserName = session.UserName
			session.Cart = cart
			t.Execute(w, session)
		} else {
			t := template.Must(template.ParseFiles("views/pages/cart/cart.html"))
			t.Execute(w, session)
		}
	} else {
		Login(w, r)
	}
}

//DeleteCart
func DeleteCart(w http.ResponseWriter, r *http.Request) {
	cartID := r.FormValue("cartId")

	DBdao.DeleteCartByCartID(cartID)
	GetCartInfo(w, r)
}
func DeleteCartItemByID(w http.ResponseWriter, r *http.Request) {
	cartItemID := r.FormValue("cartItemId")
	iCartItemID, _ := strconv.ParseInt(cartItemID, 10, 64)
	_, session := DBdao.IsLogin(r)
	userID := session.UserID
	cart, _ := DBdao.GetCartByUserId(userID)

	cartItems := cart.CartItems

	//删除切片中的数据
	for k, v := range cartItems {
		if v.CartItemID == iCartItemID {
			cartItems = append(cartItems[:k], cartItems[k+1:]...)
			cart.CartItems = cartItems
			//删除数据库中的数据
			DBdao.DeleteCartItemByID(cartItemID)
		}
	}
	//更新数据
	DBdao.UpdateCart(cart)
	GetCartInfo(w, r)
}

func UpdateCartItemByID(w http.ResponseWriter, r *http.Request) {
	cartItemID := r.FormValue("cartItemId")
	bookCount := r.FormValue("bookCount")
	iCartItemID, _ := strconv.ParseInt(cartItemID, 10, 64)
	iBookCount, _ := strconv.ParseInt(bookCount, 10, 64)

	_, session := DBdao.IsLogin(r)
	userID := session.UserID
	cart, _ := DBdao.GetCartByUserId(userID)

	cartItems := cart.CartItems

	//删除切片中的数据
	for _, v := range cartItems {
		if v.CartItemID == iCartItemID {
			v.Count = iBookCount
			//删除数据库中的数据
			DBdao.UpdateBookCount(v)
		}
	}
	//更新数据
	DBdao.UpdateCart(cart)
	GetCartInfo(w, r)
}
