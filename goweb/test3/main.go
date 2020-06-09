package main

import (
	"html/template"
	"net/http"
)

//创建处理器函数
func testTemplate(w http.ResponseWriter, r *http.Request) {
	//解析模板
	//t, _ := template.ParseFiles("index.html")
	//执行模板
	//t.Execute(w, "hello")
	t := template.Must(template.ParseFiles("index.html", "index2.html"))
	t.ExecuteTemplate(w, "index2.html", "你猜")
}

func main() {
	http.HandleFunc("/testTemplate", testTemplate)
	http.ListenAndServe(":8080", nil)
}