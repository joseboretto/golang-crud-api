package book

import (
	"github.com/joseboretto/golang-crud-api/internal/domain/model"
)

type GetAllBooksServiceInterface interface {
	GetAllBooks() ([]*model.Book, error)
}

type GetAllBooksRepositoryInterface interface {
	SelectAllBooks() ([]*model.Book, error)
}
