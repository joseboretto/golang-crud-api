package book

import "github.com/joseboretto/golang-crud-api/internal/domain/model"

type GetAllBooksRepository struct {
	database InMemoryKeyValueStorage
}

func NewGetAllBooksRepository(database InMemoryKeyValueStorage) *GetAllBooksRepository {
	return &GetAllBooksRepository{
		database: database,
	}
}

func (c *GetAllBooksRepository) SelectAllBooks() ([]*model.Book, error) {
	return c.database.GetAll(), nil
}
