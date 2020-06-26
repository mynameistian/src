package main

import (
	"goweb/bookstore/controller"
	"net/http"
	_ "text/template"
)

func main() {
	//设置静态文件
	http.Handle("/static/", http.StripPrefix("/static/", http.FileServer(http.Dir("views/static"))))

	http.Handle("/pages/", http.StripPrefix("/pages/", http.FileServer(http.Dir("views/pages"))))
	//直接去访问页面
	//http.Handle("/index/", http.StripPrefix("/pages/", http.FileServer(http.Dir("views/pages"))))
	//引擎
	http.HandleFunc("/main", controller.ToUpdateBookPageByPrice)

	//登录函数
	http.HandleFunc("/login", controller.Login)
	//注销
	http.HandleFunc("/logout", controller.Logout)
	//注册
	http.HandleFunc("/regist", controller.Regist)
	//通过Ajax请求验证用户名是否可用
	http.HandleFunc("/checkUserName", controller.CheckUserName)
	//获取所有图书
	http.HandleFunc("/getBooks", controller.GetBooks)
	// //添加图书
	// http.HandleFunc("/addBook", controller.AddBook)
	//删除图书
	http.HandleFunc("/deleteBook", controller.DeleteBook)
	//查询带页数
	http.HandleFunc("/toUpdateBookPage", controller.ToUpdateBookPage)
	//查询带价格
	http.HandleFunc("/toUpdateBookPagebyPrice", controller.ToUpdateBookPageByPrice)
	//更新图书
	http.HandleFunc("/updateBook", controller.UpdateBook)
	//获取分页书列表
	http.HandleFunc("/getPageBooks", controller.GetPageBooks)
	//添加图书到购物车
	http.HandleFunc("/addBook2Cart", controller.AddBook2Cart)
	//展示购物车信息
	http.HandleFunc("/getCartInfo", controller.GetCartInfo)
	//清空购物车
	http.HandleFunc("/deleteCart", controller.DeleteCart)
	//删除购物项
	http.HandleFunc("/deleteCartItem", controller.DeleteCartItemByID)

	//更新购物项
	http.HandleFunc("/updateCartItem", controller.UpdateCartItemByID)
	//结账
	http.HandleFunc("/checkout", controller.Checkout)

	http.ListenAndServe(":8080", nil)
}
