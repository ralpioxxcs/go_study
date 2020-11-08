package main

import (
	"fmt"
	"github.com/ralpioxxcs/go_study/cgo/clib"
	"github.com/ralpioxxcs/go_study/cgo/mylogger"
	"html/template"
	"net/http"
	"strconv"
)

func mainpage(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path != "/" {
		http.Error(w, "404 not found.", http.StatusNotFound)
		return
	}

	t, _ := template.ParseFiles("static/index.html")
	switch r.Method {
	case "GET":
	case "POST":
		if err := r.ParseForm(); err != nil {
			fmt.Fprintf(w, "ParseForm() err: %v", err)
			return
		}
		a, _ := strconv.Atoi(r.FormValue("a"))
		b, _ := strconv.Atoi(r.FormValue("b"))
		fmt.Printf("a + b = %d\n", clib.Add_cpp(a, b))
	}
	t.Execute(w, nil)
}

func create(w http.ResponseWriter, r *http.Request) {
	fmt.Println("crate button")
}

func start(w http.ResponseWriter, r *http.Request) {
	fmt.Println("start button")
}

func main() {
	logger := mylogger.GetInstance()
	logger.Println("Starting")

	logger.Println("load cpp library")
	clib.LoadLib()

	http.HandleFunc("/", mainpage)
	http.HandleFunc("/create", create)
	http.HandleFunc("/start", start)

	logger.Println("http server listen")
	http.ListenAndServe(":8080", nil)

}
