package main

import (
	"fmt"
	"net/http"
)

func setCookie(w http.ResponseWriter, r *http.Request) {
	c1 := http.Cookie{
		Name:     "first_co",
		Value:    "Go web",
		HttpOnly: true,
	}
	c2 := http.Cookie{
		Name:     "secount_co",
		Value:    "Manning Publications Go",
		HttpOnly: true,
	}
	// w.Header().Set("Set-Cookie", c1.String())
	// w.Header().Add("Set-Cookie", c2.String())
	http.SetCookie(w, &c1)
	http.SetCookie(w, &c2)
}

func getCookie(w http.ResponseWriter, r *http.Request) {
	// h := r.Header["Cookie"]
	// fmt.Fprintln(w, h)
	c1, err := r.Cookie("first_co")
	if err != nil {
		fmt.Fprintln(w, "not found first cookie")
	}
	cs := r.Cookies()
	fmt.Fprintln(w, c1)
	fmt.Fprintln(w, cs)
}

func main() {
	server := http.Server{
		Addr: "localhost:8080",
	}
	http.HandleFunc("/set_cookie", setCookie)
	http.HandleFunc("/get_cookie", getCookie)
	server.ListenAndServe()
}
