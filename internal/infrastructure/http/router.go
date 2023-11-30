package http

import (
	"github.com/joseboretto/golang-crud-api/internal/infrastructure/http/controller/book"
	"net/http"
)

func SetupRoutes(controller *book.BookController) {
	http.HandleFunc("/hello", controller.GetHelloWorld)
	http.HandleFunc("/api/v1/createBook", controller.CreateBook)
	http.HandleFunc("/api/v1/getBook", controller.GetBooks)
}
