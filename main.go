package main

import (
	"go-crud-simple-cicd/config"
	"go-crud-simple-cicd/controller"
	"go-crud-simple-cicd/db"
	"go-crud-simple-cicd/repo"
	"go-crud-simple-cicd/usecase"
	"log"

	"github.com/gin-gonic/gin"
)

func main() {
	loadConfig, err := config.LoadConfig(".")
	if err != nil {
		log.Fatal("🚀 Could not load environment variables", err)
	}

	db := db.InitDB(&loadConfig)

	router := gin.Default()

	bookRepo := repo.NewBookRepo(db)
	bookUseCase := usecase.NewBookUseCase(bookRepo)
	bookController := controller.NewBookController(bookUseCase)

	router.POST("/book", bookController.CreateBook)       // Create a new book
	router.GET("/book/:id", bookController.GetBookByID)   // Get a book by ID
	router.GET("/books", bookController.GetAllBooks)      // Get all books
	router.PUT("/book/:id", bookController.UpdateBook)    // Update a book by ID
	router.DELETE("/book/:id", bookController.DeleteBook) // Delete a book by ID

	router.Run(":1312")
}
