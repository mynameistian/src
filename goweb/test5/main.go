package main

import (
	"html/template"
	"net/http"
)

//创建处理器函数
func IndexHandler(w http.ResponseWriter, r *http.Request) {
	//解析模板
	//t, _ := template.ParseFiles("index.html")
	//执行模板
	//t.Execute(w, "hello")
	t := template.Must(template.ParseFiles("hello.html"))
	t.Execute(w, "")
}

func testDefine(w http.ResponseWriter, r *http.Request) {
	t := template.Must(template.ParseFiles("hello.html"))
	t.ExecuteTemplate(w, "model", "")
}

func testDefineHandler(w http.ResponseWriter, r *http.Request) {
	name := "百度"
	var t *template.Template
	if name == "百度" {
		t = template.Must(template.ParseFiles("hello.html", "define1.html"))
	} else if name == "hao123" {
		t = template.Must(template.ParseFiles("hello.html", "define2.html"))
	} else {
		t = template.Must(template.ParseFiles("hello.html"))
	}

	t.ExecuteTemplate(w, "model", "")
}

func main() {
	http.HandleFunc("/IndexHandler", IndexHandler)
	http.HandleFunc("/testDefine", testDefine)
	http.HandleFunc("/testDefineHandler", testDefineHandler)

	http.ListenAndServe(":8080", nil)
}
