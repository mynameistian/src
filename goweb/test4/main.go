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
	t := template.Must(template.ParseFiles("index.html"))
	t.Execute(w, "")
}

func main() {
	http.HandleFunc("/IndexHandler", IndexHandler)
	http.ListenAndServe(":8080", nil)
}
