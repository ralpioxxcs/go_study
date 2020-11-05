package main

import (
	"fmt"
	"github.com/ralpioxxcs/cgo/clib"
	"github.com/ralpioxxcs/cgo/mylogger"
	"html/template"
	"net/http"
)

//func test(w http.ResponseWriter, r *http.Request) {
//if r.Method != {
//t, _ := template.ParseFiles("server.html")
//t.Execute(w, nil)

//clib.Start()
//}
//}

func control(w http.ResponseWriter, r *http.Request) {
	fmt.Println("control")

	if r.Method == "GET" {
		t, _ := template.ParseFiles("server.html")
		t.Execute(w, nil)
	} else {
		r.ParseForm()
		fmt.Println("test", r.Form["Create"])
	}
}

func main() {
	logger := mylogger.GetInstance()
	logger.Println("Starting")

	logger.Println("load cpp library")
	clib.LoadLib()

	http.HandleFunc("/control", control)

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
