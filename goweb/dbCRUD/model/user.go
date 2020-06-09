package model

import (
	"fmt"
	"goweb/dbCRUD/utils"
)

type User struct {
	Id       int
	Username string
	Password string
	Email    string
}

//AddUser 有预编译的
func (user *User) AddUser() (err error) {
	//写sql语句
	sqlStr := "insert into users(username,password,email)values($1,$2,$3);"

	//预处理
	inStmt, err := utils.Db.Prepare(sqlStr)
	if err != nil {
		fmt.Println("utils.DB.Prepare err", err)
		return err
	}
	//执行
	_, err = inStmt.Exec("tianlj", "123456", "100352143@qq.com")
	if err != nil {
		fmt.Println("inStmt.Exec err")
		return err
	}
	return
}

//AddUser2 没有预编译的
func (user *User) AddUser2() (err error) {
	//写sql语句
	sqlStr := "insert into users(username,password,email)values($1,$2,$3)"

	//执行
	_, err = utils.Db.Exec(sqlStr, "tianlj2", "123456", "100352143@qq.com")
	if err != nil {
		fmt.Println("inStmt.Exec err", err)
		return err
	}
	return
}
