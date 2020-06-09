package main

import (
	"fmt"
	"net/http"
)

//handler 创建处理器函数
func handler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "你发送的请求的请求地址：", r.URL.Path)
	fmt.Fprintln(w, "你发送的请求的请求地址后查询的字符串是：", r.URL.RawQuery)
}

func main() {
	http.HandleFunc("/hello", handler)
	http.ListenAndServe(":8080", nil)
}
