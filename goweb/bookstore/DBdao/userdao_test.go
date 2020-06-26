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
	//fmt.Println("测试bookdao中的方法")
	//t.Run("测试获取所有图书", testGetBooks)
	//t.Run("测试添加图书", testAddBooks)
	//t.Run("测试删除图书", testDeleteBook)
	//t.Run("测试获取一本图书", testGetBook)
	//t.Run("测试获取一本图书", testUpdateBooks)
	//t.Run("测试获取一本图书", testGetPageBooks)
	// t.Run("测试获取一本图书", testGetPageBooksByPrice)
	//t.Run("测试添加session", testAddSession)
	// t.Run("测试添加session", testDeleteSession)
	//t.Run("测试添加session", testGetSession)
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
		Price:   88.88,
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
		Price:   88.88,
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

func testAddSession(t *testing.T) {

	session := &model.Session{
		SessionID: "13838381438",
		UserName:  "tianlj1",
		UserID:    6,
	}

	AddSession(session)
}

func testDeleteSession(t *testing.T) {

	err := DeleteSession("13838381438")

	if err != nil {
		fmt.Println(err)
	}
	fmt.Println("ok")
}
func testGetSession(t *testing.T) {

	sess, err := GetSession("a3f0ebe3-a845-499c-4f4b-863615e36e7d")

	if err != nil {
		fmt.Println(err)
	}

	fmt.Println(sess)
}

func TestCart(t *testing.T) {
	fmt.Println("购物车相关测试")
	//t.Run("测试添加购物车", testAddCart)
	// t.Run("测试添加购物车", testGetCartItemByBookID)
	// t.Run("测试添加购物车", testGetCatrtItemByCartID)
	//t.Run("测试添加购物车", testGetCartByUserId)
	// t.Run("测试添加购物车", testUpdateBookCount)
	// t.Run("测试添加购物车", testDeleteCartByCartID)
	// t.Run("测试添加购物车", testDeleteCartItemByID)

}

func testAddCart(t *testing.T) {
	book1 := &model.Book{
		ID:    1,
		Price: 27.00,
	}
	book2 := &model.Book{
		ID:    2,
		Price: 23.00,
	}

	var cartItems []*model.CartItem

	cartItem1 := &model.CartItem{
		Book:   book1,
		Count:  10,
		CartID: "66668888",
	}
	cartItems = append(cartItems, cartItem1)
	cartItem2 := &model.CartItem{
		Book:   book2,
		Count:  10,
		CartID: "66668888",
	}

	cartItems = append(cartItems, cartItem2)
	cart := &model.Cart{
		CartID:    "66668888",
		CartItems: cartItems,
		UserID:    6,
	}

	err := AddCart(cart)
	if err != nil {
		fmt.Println(err)
	}
}

func testGetCartItemByBookID(t *testing.T) {
	cartItem, err := GetCartItemByBookID("1")
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(cartItem)
}

func testGetCatrtItemByCartID(t *testing.T) {

	cartItems, err := GetCatrtItemByCartID("66668888")
	if err != nil {
		fmt.Println(err)
	}
	for k, v := range cartItems {
		fmt.Printf("第%v购物项是: %v\n", k, v)
	}
}

func testGetCartByUserId(t *testing.T) {
	cart, err := GetCartByUserId(6)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(cart)

	for _, k := range cart.CartItems {
		fmt.Println(k)
	}
}

func testUpdateBookCount(t *testing.T) {
	// err := UpdateBookCount(20, 1, "66668888")
	// if err != nil {
	// 	fmt.Println(err)
	// }
}

func testDeleteCartByCartID(t *testing.T) {
	err := DeleteCartByCartID("62badce8-2bcb-42a7-4028-536b58e7e03b")
	if err != nil {
		fmt.Println(err)
	}
}

func testDeleteCartItemByID(t *testing.T) {
	err := DeleteCartItemByID("50")
	if err != nil {
		fmt.Println(err)
	}
}

func TestOrder(t *testing.T) {
	fmt.Println("订单相关测试")
	// t.Run("测试订单详情添加", testAddOrder)
	// t.Run("测试订单详情添加", testAddOrderItem)
	t.Run("测试订单详情添加", testUpdateOrder)
}

func testAddOrderItem(t *testing.T) {
	orderItem := &model.OrderItem{
		ID:      123,
		Count:   1,
		Amount:  13,
		Title:   "你猜",
		Author:  "田利军",
		Price:   13,
		ImgPath: "nil",
		OrderId: "123",
	}
	err := AddOrderItem(orderItem)
	if err != nil {
		fmt.Println(err)
	}
}
func testGetOrderItem(t *testing.T) {
	orders, err := GetOrderItem("123")
	if err != nil {
		fmt.Println(err)
	}
	for _, v := range orders {
		fmt.Println(v)
	}
}

func testAddOrder(t *testing.T) {
	order := &model.Order{
		ID:          123,
		CreateTime:  "2020-06-26 12:01:01",
		TotalCount:  1,
		TotalAmount: 13,
		State:       0,
		Uer_id:      6,
	}
	err := AddOrder(order)
	if err != nil {
		fmt.Println(err)
	}
}

func testGetOrder(t *testing.T) {
	orders, err := GetOrder("123")
	if err != nil {
		fmt.Println(err)
	}
	for _, v := range orders {
		fmt.Println(v)
	}
}

func testUpdateOrder(t *testing.T) {
	err := UpdateOrderState(2, "123")
	if err != nil {
		fmt.Println(err)
	}
}
