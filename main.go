package main

import (
	"fmt"
	"net/http"
	"os"
	"html/template"
	"github.com/gorilla/mux"
)

func main() {
	port := os.Getenv("PORT")

	http.Handle("/", createRouter())

	fmt.Printf("Listening on port %s...\n", port)
	err := http.ListenAndServe(":" + port, nil)

	if err != nil {
		panic(err)
	}

}

func createRouter() (*mux.Router) {
	files := http.FileServer(http.Dir("./assets/"))

	router := mux.NewRouter()
	router.HandleFunc("/", index).Methods("GET")
	router.HandleFunc("/hello", hello).Methods("GET")
	router.HandleFunc("/bye", bye).Methods("GET")
	router.PathPrefix("/").Handler(files)

	return router
}

func index(res http.ResponseWriter, req *http.Request) {
	tmpl, _ := template.ParseFiles("./views/index.html")
	tmpl.Execute(res, nil)
}

func bye(res http.ResponseWriter, req *http.Request) {
	fmt.Fprintf(res, "Goodbye cruel world!")
}

func hello(res http.ResponseWriter, req *http.Request) {
	fmt.Fprintf(res, "Hello, World!")
}
