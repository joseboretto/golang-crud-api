package book

import (
	"encoding/json"
	servicebook "github.com/joseboretto/golang-crud-api/internal/application/service/book"
	"io"
	"log"
	"net/http"
)

type Controller struct {
	createBookServiceInterface  servicebook.CreateBookServiceInterface
	getAllBooksServiceInterface servicebook.GetAllBooksServiceInterface
}

func NewBookController(createBookServiceInterface servicebook.CreateBookServiceInterface,
	getAllBooksServiceInterface servicebook.GetAllBooksServiceInterface) *Controller {
	return &Controller{
		createBookServiceInterface:  createBookServiceInterface,
		getAllBooksServiceInterface: getAllBooksServiceInterface,
	}
}

func (controller *Controller) GetHelloWorld(w http.ResponseWriter, req *http.Request) {
	io.WriteString(w, "Hello, world 2!\n")
}

func (controller *Controller) CreateBook(w http.ResponseWriter, req *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	if req.Method == "POST" {
		w.WriteHeader(http.StatusOK)
		var createBookRequest CreateBookRequest
		err := json.NewDecoder(req.Body).Decode(&createBookRequest)
		if err != nil {
			log.Fatalln("There was an error decoding the req body into the struct")
		}
		// mapper
		bookDomain := MapToBookModel(createBookRequest)
		// service
		createBook, err := controller.createBookServiceInterface.CreateBook(bookDomain)
		if err != nil {
			log.Fatalln("There was an error decoding the req body into the struct")
		}
		// response
		createBookResponse := MapToCreateBookResponse(createBook)
		err = json.NewEncoder(w).Encode(&createBookResponse)
		if err != nil {
			log.Fatalln("There was an error encoding the initialized struct")
		}
	} else {
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte("Bad Request"))
	}

}

func (controller *Controller) GetBooks(w http.ResponseWriter, req *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	if req.Method == "GET" {
		w.WriteHeader(http.StatusOK)
		// service
		books, err := controller.getAllBooksServiceInterface.GetAllBooks()
		if err != nil {
			log.Fatalln("There was an error decoding the req body into the struct")
		}
		// response
		booksResponse := make([]*GetAllBookResponse, 0, len(books))
		for _, value := range books {
			bookResponse := MapToGetAllBookResponse(value)
			booksResponse = append(booksResponse, bookResponse)
		}
		err = json.NewEncoder(w).Encode(&booksResponse)
		if err != nil {
			log.Fatalln("There was an error encoding the initialized struct")
		}
	} else {
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte("Bad Request"))
	}

}
