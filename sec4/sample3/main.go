package main

import (
	"fmt"
	"net/http"
)

func process(w http.ResponseWriter, r *http.Request) {
	r.ParseForm()
	// // fmt.Fprintln(w, r.Form)
	// fmt.Fprintln(w, r.Form.Get("post"))
	fmt.Fprintln(w, r.PostForm)
}

func main() {
	server := http.Server{
		Addr: "localhost:8080",
	}
	http.HandleFunc("/process", process)
	server.ListenAndServe()
}
