package book

import "github.com/joseboretto/golang-crud-api/internal/domain/model"

type CreateBookRepository struct {
	database InMemoryBookDatabase
}

func NewCreateBookRepository(database InMemoryBookDatabase) *CreateBookRepository {
	return &CreateBookRepository{
		database: database,
	}
}

func (c *CreateBookRepository) InsertBook(book *model.Book) (*model.Book, error) {
	insert, err := c.database.Insert(book)
	if err != nil {
		return nil, err
	}
	return insert, nil
}
