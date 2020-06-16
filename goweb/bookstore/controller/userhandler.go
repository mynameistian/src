package controller

import (
	_ "fmt"
	"goweb/bookstore/DBdao"
	"goweb/bookstore/model"
	"goweb/bookstore/utils"
	"net/http"
	"text/template"
)

//logout 处理用处注销的函数
func Logout(w http.ResponseWriter, r *http.Request) {
	//获取Cookie
	cookie, _ := r.Cookie("user")
	if cookie != nil {
		cookieValue := cookie.Value
		DBdao.DeleteSession(cookieValue)
		//设置cookie失效
		cookie.MaxAge = -1
		http.SetCookie(w, cookie)
	}
	ToUpdateBookPageByPrice(w, r)
}

//Login 处理用户登录的函数
func Login(w http.ResponseWriter, r *http.Request) {

	//判断是否已经登录

	flag, _ := DBdao.IsLogin(r)
	if flag {
		ToUpdateBookPageByPrice(w, r)
	} else {
		//获取用户名和密码
		username := r.PostFormValue("username")
		password := r.PostFormValue("password")

		user, _ := DBdao.CheckUserNameAndPassword(username, password)
		if user.ID > 0 {
			//创建uuid
			uuid := utils.CreateUUID()
			//创建session
			sess := &model.Session{
				SessionID: uuid,
				UserName:  user.Username,
				UserID:    user.ID,
			}
			DBdao.AddSession(sess)
			//创建cookie
			cookie := http.Cookie{
				Name:     "user",
				Value:    uuid,
				HttpOnly: true,
			}
			//添加cookie
			http.SetCookie(w, &cookie)
			//用户名和密码正确
			t := template.Must(template.ParseFiles("views/pages/user/login_success.html"))
			t.Execute(w, user.Username)
		} else {
			//用户密码不正确
			t := template.Must(template.ParseFiles("views/pages/user/login.html"))
			t.Execute(w, "")
		}
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
