package book

import "github.com/joseboretto/golang-crud-api/internal/domain/model"

type InMemoryKeyValueStorage struct {
	// This uses values over pointers to avoid unexpected changes
	// This is not thread-safe, and it's not meant to be used in production.
	// booksMap is a map of books, using the isbn as the key.
	// In this case, the database is using the model, but it should be an Entity.
	booksMap map[string]model.Book
}

func NewInMemoryKeyValueStorage() *InMemoryKeyValueStorage {
	return &InMemoryKeyValueStorage{
		booksMap: make(map[string]model.Book),
	}
}

func (i *InMemoryKeyValueStorage) Insert(book *model.Book) (*model.Book, error) {
	i.booksMap[book.Isbn] = *book
	return book, nil
}

func (i *InMemoryKeyValueStorage) GetAll() []*model.Book {
	books := make([]*model.Book, 0, len(i.booksMap))
	for _, value := range i.booksMap {
		books = append(books, &value)
	}
	return books
}
