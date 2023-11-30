package main

import (
	servicebook "github.com/joseboretto/golang-crud-api/internal/application/service/book"
	infrastructurehttp "github.com/joseboretto/golang-crud-api/internal/infrastructure/http"
	controllerbook "github.com/joseboretto/golang-crud-api/internal/infrastructure/http/controller/book"

	persistancebook "github.com/joseboretto/golang-crud-api/internal/infrastructure/persistance/book"
	"log"
	"net/http"
)

func main() {
	// Hello world, the web server
	database := persistancebook.NewInMemoryBookDatabase()
	// repositories
	newCreateBookRepository := persistancebook.NewCreateBookRepository(*database)
	newGetAllBooksRepository := persistancebook.NewGetAllBooksRepository(*database)
	// services
	createBookService := servicebook.NewCreateBookService(newCreateBookRepository)
	getAllBooksService := servicebook.NewGetAllBooksService(newGetAllBooksRepository)
	bookController := controllerbook.NewBookController(createBookService, getAllBooksService)
	infrastructurehttp.SetupRoutes(bookController)

	log.Println("Listing for requests at http://localhost:8000/hello")
	log.Fatal(http.ListenAndServe(":8000", nil))
}
