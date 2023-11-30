package main

import (
	servicebook "github.com/joseboretto/golang-crud-api/internal/application/services/books"
	controller "github.com/joseboretto/golang-crud-api/internal/infrastructure/controllers"
	controllerbook "github.com/joseboretto/golang-crud-api/internal/infrastructure/controllers/books"
	"github.com/joseboretto/golang-crud-api/internal/infrastructure/persistance"

	persistancebook "github.com/joseboretto/golang-crud-api/internal/infrastructure/persistance/book"
	"log"
	"net/http"
)

func main() {
	// Dependency Injection
	database := persistance.NewInMemoryKeyValueStorage()
	// repositories
	newCreateBookRepository := persistancebook.NewCreateBookRepository(*database)
	newGetAllBooksRepository := persistancebook.NewGetAllBooksRepository(*database)
	// services
	createBookService := servicebook.NewCreateBookService(newCreateBookRepository)
	getAllBooksService := servicebook.NewGetAllBooksService(newGetAllBooksRepository)
	// controllers
	bookController := controllerbook.NewBookController(createBookService, getAllBooksService)
	helloController := controller.NewHelloController()
	// routes
	controller.SetupRoutes(bookController, helloController)

	log.Println("Listing for requests at http://localhost:8000")
	log.Fatal(http.ListenAndServe(":8000", nil))
}
