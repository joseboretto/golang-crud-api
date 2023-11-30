package book

import (
	"github.com/joseboretto/golang-crud-api/internal/domain/model"
)

type CreateBookServiceInterface interface {
	CreateBook(book *model.Book) (*model.Book, error)
}

type CreateBookRepositoryInterface interface {
	InsertBook(book *model.Book) (*model.Book, error)
}
