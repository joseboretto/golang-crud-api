package book

import "github.com/joseboretto/golang-crud-api/internal/domain/model"

type GetAllBooksRepository struct {
	database InMemoryBookDatabase
}

func NewGetAllBooksRepository(database InMemoryBookDatabase) *GetAllBooksRepository {
	return &GetAllBooksRepository{
		database: database,
	}
}

func (c *GetAllBooksRepository) SelectAllBooks() ([]model.Book, error) {
	return c.database.GetAll(), nil
}
