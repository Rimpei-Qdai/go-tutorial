package main

import (
	"fmt"
	"net/http"
)

func handler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "アクセスしたパス: %s", r.URL.Path)
}

func handler2(w http.ResponseWriter, r *http.Request) {
	//getクエリパラメータから値を取得
	name := r.URL.Query().Get("name")
	stra := r.URL.Query().Get("a")
	strb := r.URL.Query().Get("b")

	var a, b int
	fmt.Sscanf(stra, "%d", &a) // a に値を代入
	fmt.Sscanf(strb, "%d", &b) // b に値を代入
	c := a + b // a と b の合計を計算

	fmt.Fprint(w, "name:", name)
	fmt.Fprint(w, "a+b =", c)
}

func main() {
	http.HandleFunc("/", handler)
	http.HandleFunc("/add", handler2)
	fmt.Println("http://localhost:8080 でサーバーが起動しました")
	http.ListenAndServe(":8080", nil)
}