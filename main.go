package main

import (
	"fmt"
	"net/http"
	"os"
)

func main() {
	port := os.Getenv("PORT")
	http.HandleFunc("/", hello)
	fmt.Printf("Listening on port %s...\n", port)
	err := http.ListenAndServe(":" + port, nil)

	if err != nil {
		panic(err)
	}

}

func hello(res http.ResponseWriter, req *http.Request) {
	fmt.Fprintf(res, "Hello, World!")
}
