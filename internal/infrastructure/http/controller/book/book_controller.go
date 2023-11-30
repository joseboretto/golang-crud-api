package book

import (
	"encoding/json"
	servicebook "github.com/joseboretto/golang-crud-api/internal/application/service/book"
	"io"
	"log"
	"net/http"
)

type BookController struct {
	createBookServiceInterface  servicebook.CreateBookServiceInterface
	getAllBooksServiceInterface servicebook.GetAllBooksServiceInterface
}

func NewBookController(createBookServiceInterface servicebook.CreateBookServiceInterface,
	getAllBooksServiceInterface servicebook.GetAllBooksServiceInterface) *BookController {
	return &BookController{
		createBookServiceInterface:  createBookServiceInterface,
		getAllBooksServiceInterface: getAllBooksServiceInterface,
	}
}

func (controller *BookController) GetHelloWorld(w http.ResponseWriter, req *http.Request) {
	io.WriteString(w, "Hello, world 2!\n")
}

func (controller *BookController) CreateBook(w http.ResponseWriter, req *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	if req.Method == "POST" {
		w.WriteHeader(http.StatusOK)
		var createBookRequest CreateBookRequest
		err := json.NewDecoder(req.Body).Decode(&createBookRequest)
		if err != nil {
			log.Fatalln("There was an error decoding the req body into the struct")
		}
		// mapper
		bookDomain := MapToDomain(createBookRequest)
		// service
		createBook, err := controller.createBookServiceInterface.CreateBook(bookDomain)
		if err != nil {
			log.Fatalln("There was an error decoding the req body into the struct")
		}
		// response
		err = json.NewEncoder(w).Encode(&createBook)
		if err != nil {
			log.Fatalln("There was an error encoding the initialized struct")
		}
	} else {
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte("Bad Request"))
	}

}

func (controller *BookController) GetBooks(w http.ResponseWriter, req *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	if req.Method == "GET" {
		w.WriteHeader(http.StatusOK)
		// service
		books, err := controller.getAllBooksServiceInterface.GetAllBooks()
		if err != nil {
			log.Fatalln("There was an error decoding the req body into the struct")
		}
		// response
		err = json.NewEncoder(w).Encode(&books)
		if err != nil {
			log.Fatalln("There was an error encoding the initialized struct")
		}
	} else {
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte("Bad Request"))
	}

}
