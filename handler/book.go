package handler

import (
	"fmt"
	"net/http"
	"pustaka-api/book"

	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
)

func RootHandler(ctx *gin.Context) {
	ctx.JSON(http.StatusOK, gin.H{
		"name": "Herman",
	})
}

func BooksHandler(ctx *gin.Context) {
	id := ctx.Param("id")
	ctx.JSON(http.StatusOK, gin.H{"id": id})
}

func QueryHandler(ctx *gin.Context) {
	title := ctx.Query("title")
	ctx.JSON(http.StatusOK, gin.H{"title": title})
}

func PostBooksHandler(ctx *gin.Context) {
	var bookInput book.BookInput

	err := ctx.ShouldBindJSON(&bookInput)
	if err != nil {
		// log.Fatal(err)
		errorMessages := []string{}

		for _, e := range err.(validator.ValidationErrors) {
			errorMessage := fmt.Sprintf("Error on field %s, condition: %s", e.Field(), e.ActualTag())
			errorMessages = append(errorMessages, errorMessage)
		}

		ctx.JSON(http.StatusBadRequest, gin.H{
			"errors": errorMessages,
		})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{
		"title": bookInput.Title,
		"price": bookInput.Price,
		// "subtitle": bookInput.Subtitle,
	})
}