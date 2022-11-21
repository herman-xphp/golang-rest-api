package main

import (
	"fmt"
	"log"
	"pustaka-api/book"
	"pustaka-api/handler"

	"github.com/gin-gonic/gin"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func main() {
	dsn := "root:root@tcp(127.0.0.1:3306)/pustaka-api?charset=utf8mb4&parseTime=True&loc=Local"
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatal("DB Connection Error")
	}

	db.AutoMigrate(&book.Book{})
	fmt.Println("AutoMigrate Success")

	// CRUD

	// Create
	book := book.Book{}
	book.Title = "Man Tiger"
	book.Price = 90000
	book.Description = "ini adalah buku bagus"
	book.Rating = 5

	err = db.Create(&book).Error
	if err != nil {
		fmt.Println("================================")
		fmt.Println("Error creating book record")
		fmt.Println("================================")
	}

	

	router := gin.Default()

	v1 := router.Group("/v1")

	v1.GET("/", handler.RootHandler)
	v1.GET("/books/:id", handler.BooksHandler)
	v1.GET("/query", handler.QueryHandler)
	v1.POST("/books", handler.PostBooksHandler)

	router.Run(":8000")
}
