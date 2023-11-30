package book

import "github.com/joseboretto/golang-crud-api/internal/domain/model"

type InMemoryBookDatabase struct {
	books []model.Book
}

func NewInMemoryBookDatabase() *InMemoryBookDatabase {
	return &InMemoryBookDatabase{}
}

func (i *InMemoryBookDatabase) Insert(book *model.Book) (*model.Book, error) {
	i.books = append(i.books, *book)
	return book, nil
}

func (i *InMemoryBookDatabase) GetAll() []model.Book {
	return i.books
}
