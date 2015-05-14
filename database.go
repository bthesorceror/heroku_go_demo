package main

import (
	"os"
	"github.com/jinzhu/gorm"
	_ "github.com/lib/pq"
)


type Post struct {
	ID int64
	Title string
	Body string
	Owner string
	Type string
}

func newPost(title string, body string) (Post) {
	return Post{
		Title: title,
		Body: body,
	}
}

func InitDB() (gorm.DB) {
	url := os.Getenv("DATABASE_URL")
	db, err := gorm.Open("postgres", url)

	if err != nil {
		panic(err)
	}

	db.AutoMigrate(&Post{})

	return db
}
