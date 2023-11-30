package controller

import (
	"github.com/joseboretto/golang-crud-api/internal/infrastructure/controller/book"
	"net/http"
)

func SetupRoutes(bookController *book.Controller, helloController *HelloController) {
	http.HandleFunc("/", helloController.HelloWorld)
	http.HandleFunc("/api/v1/createBook", bookController.CreateBook)
	http.HandleFunc("/api/v1/getBook", bookController.GetBooks)
}
