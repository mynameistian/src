package main

import (
	"encoding/json"
	"fmt"
	"net/http"
)

type User struct {
	ID       int    `json:"id"`
	Username string `json:"username"`
	Password string `json:"password"`
	Email    string `json:"email"`
}

//handler 创建处理器函数
func handler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "你发送的请求的请求地址：", r.URL.Path)
	fmt.Fprintln(w, "你发送的请求的请求地址后查询的字符串是：", r.URL.RawQuery)
	fmt.Fprintln(w, "请求的所有信息有：", r.Header)
	fmt.Fprintln(w, "请求的Accept-Encoding信息是：", r.Header["Accept-Encoding"])
	fmt.Fprintln(w, "请求的Accept-Encoding信息是：", r.Header.Get("Accept-Encoding"))
	//获取请求体内内容长度
	//len := r.ContentLength
	//创建切片
	//body := make([]byte, len)
	//将请求体中的内容读到body中
	//r.Body.Read(body)
	//在浏览器中显示的请求体中的内容
	//fmt.Fprintln(w, "请求体中的内容:", string(body))

	//解析表单, 在调用r.From之前必须执行解析操作
	//r.ParseForm()
	//获取解析参数
	//如果form表单的action属性的地址中也有与form表单参数名相同参数名相同的请求参数
	//那么参数值表单中的值在url值 前面
	fmt.Fprintln(w, "请求参数有：", r.PostFormValue("username"))

}

//testJsonRes
func testJsonRes(w http.ResponseWriter, r *http.Request) {
	//设置相应类型
	r.Header.Set("Content-Type", "application/json")
	//创建User
	user := User{
		ID:       1,
		Username: "admin",
		Password: "123456",
		Email:    "admin@",
	}

	data, err := json.Marshal(user)
	if err != nil {
		fmt.Println("Marshal err is ", err)
	}
	w.Write(data)

}

func testRedire(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Location", "https://www.baidu.com/")
	w.WriteHeader(302)
}

func main() {
	http.HandleFunc("/hello", handler)
	http.HandleFunc("/json", testJsonRes)
	http.HandleFunc("/Redire", testRedire)
	http.ListenAndServe(":8080", nil)
}
