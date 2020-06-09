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
	http.HandleFunc("/main", controller.IndexHandler)

	//登录函数
	http.HandleFunc("/login", controller.Login)
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
	http.ListenAndServe(":8080", nil)
}
