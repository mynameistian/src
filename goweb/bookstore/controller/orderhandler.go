package controller

import (
	_ "fmt"
	"goweb/bookstore/DBdao"
	"goweb/bookstore/model"
	"goweb/bookstore/utils"
	"html/template"
	"net/http"
	_ "strconv"
	"time"
)

func Checkout(w http.ResponseWriter, r *http.Request) {

	var userID int
	bLogin, session := DBdao.IsLogin(r)
	if bLogin {
		userID = session.UserID
		//将用户购物车中的内容添加到订单中
		cart, _ := DBdao.GetCartByUserId(userID)

		CreatTime := time.Now().String()
		orderId := utils.CreateUUID()
		if cart != nil {
			//将购物车添加订单信息中
			order := &model.Order{
				ID:          orderId,
				CreateTime:  CreatTime,
				TotalAmount: cart.TotalAmount,
				TotalCount:  cart.TotalCount,
				State:       0,
				Uer_id:      int64(userID),
			}
			DBdao.AddOrder(order)
			session.Order = order
			cartItems, _ := DBdao.GetCatrtItemByCartID(cart.CartID)
			if cartItems != nil {
				var orderItems []*model.OrderItem
				for _, v := range cartItems {
					orderItem := &model.OrderItem{
						Count:   v.Count,
						Amount:  v.Amount,
						Title:   v.Book.Title,
						Author:  v.Book.Author,
						Price:   v.Book.Price,
						ImgPath: v.Book.ImgPath,
						OrderId: orderId,
					}
					DBdao.AddOrderItem(orderItem)
					orderItems = append(orderItems, orderItem)
				}
				//删除购物项内的数据
				DBdao.DeleteCartItemsByCartID(cart.CartID)
			} else {
				//数据存在问题，购物车中没有详细购物项
			}
			//清空购物车数据
			DBdao.DeleteCartByCartID(cart.CartID)
		}
		t := template.Must(template.ParseFiles("views/pages/cart/checkout.html"))
		t.Execute(w, session)
		//跳转页面
	} else {

	}
}
