package main

import (
	"fmt"
	"net/http"
)

func handler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "アクセスしたパス: %s", r.URL.Path)
}

func main() {
	http.HandleFunc("/", handler)
	fmt.Println("http://localhost:8080 でサーバーが起動しました")
	http.ListenAndServe(":8080", nil)
}