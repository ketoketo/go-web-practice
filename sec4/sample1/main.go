package main

import (
	"fmt"
	"net/http"
)


func headers(w http.ResponseWriter, r *http.Request) {
	// h := r.Header
	// h := r.Header["Accept-Encoding"]	// 文字列のマップ
	h := r.Header.Get("Accept-Encoding")	// カンマ区切りのリスト
	fmt.Fprintln(w, h)
}
//Zenkakuhankaku
func main() {
	server := http.Server{
		Addr: "localhost:8080",
	}
	http.HandleFunc("/headers", headers)
	server.ListenAndServeTLS("../../sec3/tls/server.crt", "../../sec3/tls/server.key")
}