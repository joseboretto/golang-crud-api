package books

import "github.com/joseboretto/golang-crud-api/internal/domain/models"

type CreateBookServiceInterface interface {
	CreateBook(book *models.Book) (*models.Book, error)
}

type CreateBookRepositoryInterface interface {
	InsertBook(book *models.Book) (*models.Book, error)
}
