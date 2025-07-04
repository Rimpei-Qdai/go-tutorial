package main

import (
	"fmt"
	"net/http"
	"html/template"
)

func handler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "アクセスしたパス: %s <html><h1>見出しです</h1></html>", r.URL.Path)
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

	tmpls := template.Must(template.New("index").Parse(`
	<!DOCTYPE html>
	<html>
	<head><title>計算結果</title></head>
	<body>
		<h1>計算結果</h1>
		<p>name: {{.Name}}</p>
		<p>a + b = {{.Result}}</p>
	</body>
	</html>			
	`))

	data := struct {
		Name   string
		Result int
	}{
		Name:   name,
		Result: c,
	}
	tmpls.Execute(w, data)
}

func main() {
	http.HandleFunc("/", handler)
	http.HandleFunc("/add", handler2)
	fmt.Println("http://localhost:8080 でサーバーが起動しました")
	http.ListenAndServe(":8080", nil)
}