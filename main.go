package main

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
)

func main() {
	// Create Object Gin Router
	router := gin.Default()

	// Create API Path /
	// router.GET("/", func(ctx *gin.Context) {
	// 	ctx.JSON(http.StatusOK, "Hello World!")
	// })

	// Creat API Bio
	// router.GET("/bio", func(ctx *gin.Context) {
	// 	ctx.JSON(http.StatusOK, gin.H{
	// 		"name":    "Ravi Mukti",
	// 		"city":    "Bandung",
	// 		"message": "Hello World From Gin!",
	// 	})
	// })

	router.GET("/", rootHandler)
	router.GET("/bio", bioHandler)
	// Using Path Variable
	router.GET("/books/:id", bookHandler)
	// Using Query Parameter
	router.GET("/books/search", queryHandler)
	// Using Request Body
	router.POST("/books", createBookHandler)

	// Run Server
	router.Run(":8000")
}

// Create RootHandler
func rootHandler(ctx *gin.Context) {
	ctx.JSON(http.StatusOK, "Hello World!")
}

// Create BioHandler
func bioHandler(ctx *gin.Context) {
	ctx.JSON(http.StatusOK, gin.H{
		"name":    "Ravi Mukti",
		"city":    "Bandung",
		"message": "Hello World From Gin!",
	})
}

// BookHandler
func bookHandler(ctx *gin.Context) {
	// Get Path Variable using Param
	id := ctx.Param("id")

	ctx.JSON(http.StatusOK, gin.H{
		"id": id,
	})
}

// QueryHandler
func queryHandler(ctx *gin.Context) {
	// Get Query Param
	title := ctx.Query("title")

	ctx.JSON(http.StatusOK, gin.H{
		"title": title,
	})
}

type BookInput struct {
	Title       string `binding:"required" json:"title"`
	Price       int    `binding:"required,number" json:"price"`
	PublishYear int    `binding:"required,number" json:"publish_year"`
}

// CreateBookHandler
func createBookHandler(ctx *gin.Context) {
	var bookInput BookInput

	err := ctx.ShouldBindJSON(&bookInput)

	if err != nil {

		errorMessages := []string{}

		for _, e := range err.(validator.ValidationErrors) {
			errorMessage := fmt.Sprintf("Error On Field %s, condition: %s", e.Field(), e.ActualTag())
			errorMessages = append(errorMessages, errorMessage)
		}
		ctx.JSON(http.StatusBadRequest, gin.H{
			"error_detail": errorMessages,
		})
		return
	}

	ctx.JSON(http.StatusCreated, gin.H{
		"title":        bookInput.Title,
		"price":        bookInput.Price,
		"publish_year": bookInput.PublishYear,
	})
}
