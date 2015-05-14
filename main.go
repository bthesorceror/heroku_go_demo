package main

import (
	"fmt"
	"net/http"
	"os"
	"html/template"
	"github.com/jinzhu/gorm"
	"github.com/gorilla/mux"
)

type IndexHandler struct {
	DB gorm.DB
}

func (h *IndexHandler) ServeHTTP(res http.ResponseWriter, req *http.Request) {
	var posts []Post
	h.DB.Find(&posts)
	tmpl, _ := template.ParseFiles("./views/index.html")
	tmpl.Execute(res, posts)
}

func main() {
	port := os.Getenv("PORT")

	http.Handle("/", createRouter())

	fmt.Printf("Listening on port %s...\n", port)
	err := http.ListenAndServe(":" + port, nil)

	if err != nil {
		panic(err)
	}
}

func setupPosts(db gorm.DB) {
	db.Delete(Post{})

	post1 := newPost("Test", "Testing Post!")
	post2 := newPost("Awesome Sauce", "That is what I am talking about!!")

	db.Create(&post1)
	db.Create(&post2)
}

func createRouter() (*mux.Router) {
	db := InitDB()
	setupPosts(db)

	files := http.FileServer(http.Dir("./assets/"))

	router := mux.NewRouter()
	router.Handle("/", &IndexHandler{DB: db }).Methods("GET")
	router.HandleFunc("/hello", hello).Methods("GET")
	router.HandleFunc("/bye", bye).Methods("GET")
	router.PathPrefix("/").Handler(files)

	return router
}

func bye(res http.ResponseWriter, req *http.Request) {
	fmt.Fprintf(res, "Goodbye cruel world!")
}

func hello(res http.ResponseWriter, req *http.Request) {
	fmt.Fprintf(res, "Hello, World!")
}
