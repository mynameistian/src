package controller

import (
	_ "fmt"
	"goweb/bookstore/DBdao"
	"net/http"
	"text/template"
)

//Login 处理用户登录的函数
func Login(w http.ResponseWriter, r *http.Request) {
	//获取用户名和密码
	username := r.PostFormValue("username")
	password := r.PostFormValue("password")

	user, _ := DBdao.CheckUserNameAndPassword(username, password)
	if user.ID > 0 {
		//用户名和密码正确
		t := template.Must(template.ParseFiles("views/pages/user/login_success.html"))
		t.Execute(w, "")
	} else {
		//用户密码不正确
		t := template.Must(template.ParseFiles("views/pages/user/login.html"))
		t.Execute(w, "")
	}
}

//Regist 处理用户登录的函数
func Regist(w http.ResponseWriter, r *http.Request) {
	//获取用户名和密码
	username := r.PostFormValue("username")
	password := r.PostFormValue("password")
	email := r.PostFormValue("email")

	user, _ := DBdao.CheckUserName(username)
	if user.ID > 0 {
		//用户名已存在
		t := template.Must(template.ParseFiles("views/pages/user/regist.html"))
		t.Execute(w, "用户名已存在!")
	} else {
		//用户名可用
		err := DBdao.SaveUser(username, password, email)
		if err != nil {
			t := template.Must(template.ParseFiles("views/pages/user/regist.html"))
			t.Execute(w, "用户名可用！")
		} else {
			t := template.Must(template.ParseFiles("views/pages/user/regist_success.html"))
			t.Execute(w, "注册失败！")
		}
	}
}

func CheckUserName(w http.ResponseWriter, r *http.Request) {
	username := r.PostFormValue("username")
	user, _ := DBdao.CheckUserName(username)
	if user.ID > 0 {
		w.Write([]byte("用户名已存在！"))
	} else {
		w.Write([]byte("<font style='color:green'>用户名可用！"))
	}
}
