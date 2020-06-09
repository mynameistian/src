package DBdao

import (
	_ "fmt"
	"goweb/bookstore/model"
	"goweb/bookstore/utils"
	"strconv"
)

//GetBooks 获取数据库中所有的图书
func GetBooks() ([]*model.Book, error) {
	//写sql语句
	sqlStr := "select id ,title, author, price,sales,stock,img_path from books"

	//执行
	rows, err := utils.Db.Query(sqlStr)
	if err != nil {
		return nil, err
	}

	var books []*model.Book
	for rows.Next() {
		book := &model.Book{}
		rows.Scan(&book.ID, &book.Title, &book.Author, &book.Price, &book.Sales, &book.Stock, &book.ImgPath)
		// fmt.Printf("ID %v \t Tiltle %v \t Author %v \t Price %v \t Sales %v \t Stock %v \t Imgpath %v \t \n", book.ID, book.Title, book.Author, book.Price, book.Sales, book.Stock, book.ImgPath)
		books = append(books, book)
	}
	return books, nil
}

//AddBooks 向数据库添加一本图书
func AddBooks(b *model.Book) error {
	sqlStr := "insert into books(title , author ,price ,sales ,stock,img_path)values($1,$2,$3,$4,$5,$6);"

	_, err := utils.Db.Exec(sqlStr, b.Title, b.Author, b.Price, b.Sales, b.Stock, b.ImgPath)
	if err != nil {
		return err
	}
	return nil
}

//DeleteBook 删除数据库中的书
func DeleteBook(bookId int) error {
	sqlStr := "delete from books where id = $1;"

	_, err := utils.Db.Exec(sqlStr, bookId)
	if err != nil {
		return err
	}
	return nil
}

//GetBookByID 根据图书的id 从数据库中查询出一本图书
func GetBookByIdID(bookID int) (*model.Book, error) {

	sqlStr := "select id ,title, author, price,sales,stock,img_path from books where id = $1;"

	row := utils.Db.QueryRow(sqlStr, bookID)
	book := &model.Book{}
	row.Scan(&book.ID, &book.Title, &book.Author, &book.Price, &book.Sales, &book.Stock, &book.ImgPath)

	return book, nil
}

//UpdateBook 更新一本图书
func UpdateBook(b *model.Book) error {

	sqlStr := "update books set title = $1 , author = $2 ,price = $3 ,sales = $4 ,stock = $5,img_path = $6 where id = $7;"

	_, err := utils.Db.Exec(sqlStr, b.Title, b.Author, b.Price, b.Sales, b.Stock, b.ImgPath, b.ID)
	if err != nil {
		return err
	}
	return nil
}

//GetPageBooks 获取带分页的图书信息
func GetPageBooks(pageNo string) (*model.Page, error) {

	//将页码转换为int64
	iPageNo, _ := strconv.ParseInt(pageNo, 10, 64)
	//获取总数
	sqlStr := "select count(*) from books;"

	var totalRecord int64
	var pageSize int64
	var TotalPageNo int64
	row := utils.Db.QueryRow(sqlStr)
	row.Scan(&totalRecord)

	pageSize = 4

	if totalRecord%pageSize == 0 {
		TotalPageNo = totalRecord / 4
	} else {
		TotalPageNo = totalRecord/4 + 1
	}

	sqlStr1 := "select id ,title, author, price,sales,stock,img_path from books limit $1 offset $2; "
	rows, err := utils.Db.Query(sqlStr1, pageSize, (iPageNo-1)*pageSize)
	if err != nil {
		return nil, err
	}
	var books []*model.Book
	for rows.Next() {
		book := &model.Book{}
		rows.Scan(&book.ID, &book.Title, &book.Author, &book.Price, &book.Sales, &book.Stock, &book.ImgPath)

		books = append(books, book)
	}

	page := &model.Page{
		Books:       books,
		PageNo:      iPageNo,
		PageSize:    pageSize,
		TotalPageNo: TotalPageNo,
		TotalRecord: totalRecord,
	}

	return page, nil
}

func GetPageBooksByPrice(pageNo string, min string, max string) (*model.Page, error) {

	//将页码转换为int64
	iPageNo, _ := strconv.ParseInt(pageNo, 10, 64)
	imin, _ := strconv.ParseInt(min, 10, 64)
	imax, _ := strconv.ParseInt(max, 10, 64)
	//获取总数
	sqlStr := "select count(*) from books where price BETWEEN $1 and $2;"

	var totalRecord int64
	var pageSize int64
	var TotalPageNo int64
	row := utils.Db.QueryRow(sqlStr, imin, imax)
	row.Scan(&totalRecord)

	pageSize = 4

	if totalRecord%pageSize == 0 {
		TotalPageNo = totalRecord / 4
	} else {
		TotalPageNo = totalRecord/4 + 1
	}

	sqlStr1 := "select id ,title, author, price,sales,stock,img_path from books where price BETWEEN $1 and $2 limit $3 offset $4; "
	rows, err := utils.Db.Query(sqlStr1, imin, imax, pageSize, (iPageNo-1)*pageSize)
	if err != nil {
		return nil, err
	}

	var books []*model.Book
	for rows.Next() {
		book := &model.Book{}
		rows.Scan(&book.ID, &book.Title, &book.Author, &book.Price, &book.Sales, &book.Stock, &book.ImgPath)

		books = append(books, book)
	}

	page := &model.Page{
		Books:       books,
		PageNo:      iPageNo,
		PageSize:    pageSize,
		TotalPageNo: TotalPageNo,
		TotalRecord: totalRecord,
	}
	return page, nil
}
