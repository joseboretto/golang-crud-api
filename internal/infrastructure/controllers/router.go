package controllers

import (
	"github.com/joseboretto/golang-crud-api/internal/infrastructure/controllers/books"
	"net/http"
)

func SetupRoutes(bookController *books.Controller, helloController *HelloController) {
	http.HandleFunc("/", helloController.HelloWorld)
	http.HandleFunc("/api/v1/createBook", bookController.CreateBook)
	http.HandleFunc("/api/v1/getBook", bookController.GetBooks)
}
