package main

import (
	"fmt"
	"github.com/ralpioxxcs/cgo/clib"
	"github.com/ralpioxxcs/cgo/mylogger"
	"html/template"
	"net/http"
)

func mainpage(w http.ResponseWriter, r *http.Request) {
	t, _ := template.ParseFiles("static/index.html")
	t.Execute(w, nil)
}

func create(w http.ResponseWriter, r *http.Request) {
	fmt.Println("crate")
	t, _ := template.ParseFiles("static/index.html")
	t.Execute(w, nil)
}

func start(w http.ResponseWriter, r *http.Request) {
	fmt.Println("start")
	t, _ := template.ParseFiles("static/index.html")
	t.Execute(w, nil)
	//if r.Method == "GET" {
	//  fmt.Println("test", r.Form["Create"])
	//  t, _ := template.ParseFiles("static/index.html")
	//  t.Execute(w, nil)
	//} else {
	//  r.ParseForm()
	//  fmt.Println("test", r.Form["Create"])
	//}
}

func main() {
	logger := mylogger.GetInstance()
	logger.Println("Starting")

	logger.Println("load cpp library")
	clib.LoadLib()

	http.HandleFunc("/", mainpage)
	http.HandleFunc("/create/", create)
	http.HandleFunc("/create", create)
	http.HandleFunc("/start/", start)
	http.HandleFunc("/start", start)

	// tmpl := template.Must(template.ParseFiles("server.html"))
	// http.HandleFunc("/", func(res http.ResponseWriter, req *http.Request) {
	// 	if req.Method != http.MethodPost {
	// 		tmpl.Execute(res, nil)
	// 		return
	// 	}
	//})

	//http.HandleFunc("/", test)
	logger.Println("http server listen")
	http.ListenAndServe(":8080", nil)

}
