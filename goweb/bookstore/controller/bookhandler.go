package controller

import (
	_ "fmt"
	"goweb/bookstore/DBdao"
	"goweb/bookstore/model"
	"html/template"
	"net/http"
	"strconv"
)

//GetBooks 获取所有图书
func GetBooks(w http.ResponseWriter, r *http.Request) {
	//调用bookdao中获取所有图书的函数
	books, _ := DBdao.GetBooks()
	t := template.Must(template.ParseFiles("views/pages/manager/book_manager.html"))
	t.Execute(w, books)
}

// //添加图书
// func AddBook(w http.ResponseWriter, r *http.Request) {
// 	title := r.PostFormValue("title")
// 	author := r.PostFormValue("author")
// 	price := r.PostFormValue("price")
// 	sales := r.PostFormValue("sales")
// 	stock := r.PostFormValue("stock")

// 	iSales, _ := strconv.Atoi(sales)
// 	iStock, _ := strconv.Atoi(stock)
// 	book := &model.Book{
// 		Title:  title,
// 		Author: author,
// 		Price:  price,
// 		Sales:  iSales,
// 		Stock:  iStock,
// 	}

// 	DBdao.AddBooks(book)
// 	//调用bookdao中获取所有图书的函数
// 	GetBooks(w, r)
// }

func DeleteBook(w http.ResponseWriter, r *http.Request) {
	bookId := r.FormValue("bookId")
	ibookId, _ := strconv.Atoi(bookId)

	// 删除图书
	DBdao.DeleteBook(ibookId)
	GetPageBooks(w, r)
}

func ToUpdateBookPage(w http.ResponseWriter, r *http.Request) {

	bookID := r.FormValue("bookId")
	ibookId, _ := strconv.Atoi(bookID)

	book, _ := DBdao.GetBookByIdID(ibookId)

	if book.ID > 0 {
		//存在更新图书
		t := template.Must(template.ParseFiles("views/pages/manager/book_edit.html"))
		t.Execute(w, book)
	} else {
		t := template.Must(template.ParseFiles("views/pages/manager/book_edit.html"))
		t.Execute(w, "")
	}

}

func UpdateBook(w http.ResponseWriter, r *http.Request) {

	id := r.PostFormValue("bookId")
	title := r.PostFormValue("title")
	author := r.PostFormValue("author")
	price := r.PostFormValue("price")
	sales := r.PostFormValue("sales")
	stock := r.PostFormValue("stock")

	iSales, _ := strconv.Atoi(sales)
	iStock, _ := strconv.Atoi(stock)
	iid, _ := strconv.Atoi(id)
	book := &model.Book{
		ID:     iid,
		Title:  title,
		Author: author,
		Price:  price,
		Sales:  iSales,
		Stock:  iStock,
	}
	if iid == 0 {
		DBdao.AddBooks(book)
	} else {
		DBdao.UpdateBook(book)
	}
	//fmt.Println("id ", id)
	//调用bookdao中获取所有图书的函数
	GetPageBooks(w, r)
}

func GetPageBooks(w http.ResponseWriter, r *http.Request) {

	page := r.FormValue("pageNo")

	if page == "" {
		page = "1"
	}
	//fmt.Println(page)
	//调用bookdao中获取所有图书的函数
	Page, _ := DBdao.GetPageBooks(page)
	t := template.Must(template.ParseFiles("views/pages/manager/book_manager.html"))
	t.Execute(w, Page)
}

//IndexHandler 去首页
func IndexHandler(w http.ResponseWriter, r *http.Request) {

	page := r.FormValue("pageNo")

	if page == "" {
		page = "1"
	}

	//调用bookdao中获取所有图书的函数

	Page, _ := DBdao.GetPageBooks(page)
	t := template.Must(template.ParseFiles("views/index.html"))
	t.Execute(w, Page)
}

//ToUpdateBookPageByPrice 带价格查询
func ToUpdateBookPageByPrice(w http.ResponseWriter, r *http.Request) {

	page := r.FormValue("pageNo")

	if page == "" {
		page = "1"
	}
	priceMin := r.FormValue("min")
	priceMax := r.FormValue("max")

	var Page *model.Page
	if priceMin == "" && priceMax == "" {
		Page, _ = DBdao.GetPageBooks(page)
	} else {
		Page, _ = DBdao.GetPageBooksByPrice(page, priceMin, priceMax)
		Page.MinPrice = priceMin
		Page.MaxPrice = priceMax
	}
	t := template.Must(template.ParseFiles("views/index.html"))
	t.Execute(w, Page)
}
