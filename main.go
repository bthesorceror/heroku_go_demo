package main

import (
	"fmt"
	"net/http"
	"os"
	"github.com/gorilla/mux"
)

func main() {
	port := os.Getenv("PORT")

	router := mux.NewRouter()
	router.HandleFunc("/", hello).Methods("GET")
	router.HandleFunc("/bye", bye).Methods("GET")
	http.Handle("/", router)

	fmt.Printf("Listening on port %s...\n", port)
	err := http.ListenAndServe(":" + port, nil)

	if err != nil {
		panic(err)
	}

}

func bye(res http.ResponseWriter, req *http.Request) {
	fmt.Fprintf(res, "Goodbye cruel world!")
}

func hello(res http.ResponseWriter, req *http.Request) {
	fmt.Fprintf(res, "Hello, World!")
}
