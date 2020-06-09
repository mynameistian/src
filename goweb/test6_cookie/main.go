package main

import (
	"fmt"
	"net/http"
)

func setCookie(w http.ResponseWriter, r *http.Request) {
	//创建Cookie

	cookie := http.Cookie{
		Name:     "user",
		Value:    "admin",
		HttpOnly: true,
		MaxAge:   60,
	}

	cookie2 := http.Cookie{
		Name:     "Name",
		Value:    "Tianlj",
		HttpOnly: true,
	}

	// w.Header().Set("Set-Cookie", cookie.String())
	// w.Header().Add("Set-Cookie", cookie2.String())

	http.SetCookie(w, &cookie)
	http.SetCookie(w, &cookie2)

}

//GetCookies
func GetCookies(w http.ResponseWriter, r *http.Request) {
	//获取所有cookies
	// cookies := r.Header["Cookie"]
	cookies, _ := r.Cookie("user")
	fmt.Println(cookies)
}

func main() {
	http.HandleFunc("/setCookie", setCookie)
	http.HandleFunc("/GetCookies", GetCookies)

	http.ListenAndServe(":8080", nil)
}
