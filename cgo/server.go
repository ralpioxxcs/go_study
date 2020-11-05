package main

import (
	"github.com/ralpioxxcs/cgo/clib"
	"github.com/ralpioxxcs/cgo/mylogger"
	"html/template"
	"net/http"
)

func main() {
	logger := mylogger.GetInstance()
	logger.Println("Starting")

	logger.Println("load cpp library")
	clib.LoadLib()

	tmpl := template.Must(template.ParseFiles("server.html"))

	http.HandleFunc("/", func(res http.ResponseWriter, req *http.Request) {
		if req.Method != http.MethodPost {
			tmpl.Execute(res, nil)
			return
		}
	})
	logger.Println("http server listen")
	http.ListenAndServe(":8080", nil)

}
