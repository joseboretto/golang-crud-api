package main

import (
	servicebook "github.com/joseboretto/golang-crud-api/internal/application/service/book"
	infrastructurehttp "github.com/joseboretto/golang-crud-api/internal/infrastructure/http"
	"github.com/joseboretto/golang-crud-api/internal/infrastructure/http/controller"
	controllerbook "github.com/joseboretto/golang-crud-api/internal/infrastructure/http/controller/book"

	persistancebook "github.com/joseboretto/golang-crud-api/internal/infrastructure/persistance/book"
	"log"
	"net/http"
)

func main() {
	// Dependency Injection
	database := persistancebook.NewInMemoryKeyValueStorage()
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
	infrastructurehttp.SetupRoutes(bookController, helloController)

	log.Println("Listing for requests at http://localhost:8000")
	log.Fatal(http.ListenAndServe(":8000", nil))
}
