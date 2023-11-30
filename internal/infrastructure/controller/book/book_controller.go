package book

import (
	servicebook "github.com/joseboretto/golang-crud-api/internal/application/service/book"
	"github.com/joseboretto/golang-crud-api/internal/infrastructure/controller/book/dto"
	"github.com/joseboretto/golang-crud-api/internal/infrastructure/controller/utils"
	"io"
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

func (c *Controller) CreateBook(w http.ResponseWriter, req *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	if req.Method == "POST" {
		createBookRequest := new(dto.CreateBookRequest)
		err := utils.Decode(req, &createBookRequest)
		if err != nil {
			w.WriteHeader(http.StatusBadRequest)
			return
		}

		// mapper
		bookDomain := dto.MapToBookModel(createBookRequest)
		// service
		createBook, err := c.createBookServiceInterface.CreateBook(bookDomain)
		if err != nil {
			w.WriteHeader(http.StatusBadRequest)
			return
		}
		// response
		createBookResponse := dto.MapToCreateBookResponse(createBook)

		if err = utils.Response(w, createBookResponse, http.StatusOK); err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			return
		}
	} else {
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte("Bad Request"))
	}

}

func (c *Controller) GetBooks(w http.ResponseWriter, req *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	if req.Method == "GET" {
		// service
		books, err := c.getAllBooksServiceInterface.GetAllBooks()
		if err != nil {
			w.WriteHeader(http.StatusBadRequest)
			return
		}
		// response
		booksResponse := make([]*dto.GetAllBookResponse, 0, len(books))
		for _, value := range books {
			bookResponse := dto.MapToGetAllBookResponse(value)
			booksResponse = append(booksResponse, bookResponse)
		}

		if err = utils.Response(w, booksResponse, http.StatusOK); err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			return
		}
	} else {
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte("Bad Request"))
	}

}
