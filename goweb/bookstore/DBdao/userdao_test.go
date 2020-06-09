package DBdao

import (
	"fmt"
	"goweb/bookstore/model"
	"testing"
)

func TestUser(t *testing.T) {
	//fmt.Println("测试userdao中的函数")
	//t.Run("验证用户名或密码", testLogin)
	//t.Run("验证用户名", testRegist)
	//t.Run("保存用户", testSave)
}

func testLogin(t *testing.T) {
	user, _ := CheckUserNameAndPassword("tianlj", "123456")
	fmt.Println(user)
}

func testRegist(t *testing.T) {
	user, _ := CheckUserName("tianlj")
	fmt.Println(user)
}
func testSave(t *testing.T) {
	SaveUser("admin2", "123456", "100352143@qq.com")
}
func TestMain(m *testing.M) {
	fmt.Println("测试bookdao中的方法")
	m.Run()
}

func TestBook(t *testing.T) {
	fmt.Println("测试bookdao中的方法")
	//t.Run("测试获取所有图书", testGetBooks)
	//t.Run("测试添加图书", testAddBooks)
	//t.Run("测试删除图书", testDeleteBook)
	//t.Run("测试获取一本图书", testGetBook)
	//t.Run("测试获取一本图书", testUpdateBooks)
	//t.Run("测试获取一本图书", testGetPageBooks)
	t.Run("测试获取一本图书", testGetPageBooksByPrice)
}

func testGetBooks(t *testing.T) {
	books, _ := GetBooks()

	for k, v := range books {
		fmt.Printf("第%v本图书的信息是: %v\n", k+1, v)
	}
}

func testAddBooks(t *testing.T) {
	book := &model.Book{
		Title:   "三国演义",
		Author:  "罗贯中",
		Price:   "88.88",
		Sales:   100,
		Stock:   100,
		ImgPath: "/station/img/default.jpg",
	}
	AddBooks(book)
}
func testDeleteBook(t *testing.T) {
	DeleteBook(65)
}

func testGetBook(t *testing.T) {
	fmt.Println(GetBookByIdID(60))
}

func testUpdateBooks(t *testing.T) {
	book := &model.Book{
		ID:      62,
		Title:   "三国演义1",
		Author:  "罗贯中",
		Price:   "88.88",
		Sales:   100,
		Stock:   100,
		ImgPath: "/station/img/default.jpg",
	}
	UpdateBook(book)
}

func testGetPageBooks(t *testing.T) {

	Page, err := GetPageBooks("16")
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(Page.PageNo)
	fmt.Println(Page.TotalPageNo)
	fmt.Println(Page.TotalRecord)
	for _, v := range Page.Books {
		fmt.Println(v)
	}

}

func testGetPageBooksByPrice(t *testing.T) {

	Page, err := GetPageBooksByPrice("1", "30", "100")
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(Page.PageNo)
	fmt.Println(Page.TotalPageNo)
	fmt.Println(Page.TotalRecord)
	for _, v := range Page.Books {
		fmt.Println(v)
	}
}
