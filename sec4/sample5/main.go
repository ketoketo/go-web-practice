package main

import (
	"encoding/json"
	"fmt"
	"net/http"
)

type Post struct {
	User    string
	Threads []string
}

func writeExample(w http.ResponseWriter, r *http.Request) {
	str := `<html>
	<head><title>Go web</title></head>
	<body><h1>Hello World</h1></body>
	</html>`
	w.Write([]byte(str))
}

func writeHeaderExample(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(501)
	fmt.Fprintln(w, "そのようなサービスはありません。")
}

func headerExample(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Location", "http://google.com")
	w.WriteHeader(302)
}

func jsonExample(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	post := &Post{
		User:    "Sau Sheong",
		Threads: []string{"no 1", "no 2", "no 3"},
	}
	json, _ := json.Marshal(post)
	w.Write(json)
}

func main() {
	server := http.Server{
		Addr: "localhost:8080",
	}
	http.HandleFunc("/write", writeExample)
	http.HandleFunc("/writeheader", writeHeaderExample)
	http.HandleFunc("/redirect", headerExample)
	http.HandleFunc("/json", jsonExample)
	server.ListenAndServe()
}
