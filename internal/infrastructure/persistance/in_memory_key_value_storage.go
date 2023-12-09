package persistance

import "github.com/joseboretto/golang-crud-api/internal/domain/models"

type InMemoryKeyValueStorage struct {
	// This uses pointers over values because it can be nil
	// This is not thread-safe, and it's not meant to be used in production.
	// booksMap is a map of books, using the isbn as the key.
	// In this case, the database is using the model, but it should be an Entity.
	booksMap map[string]*models.Book
}

func NewInMemoryKeyValueStorage() *InMemoryKeyValueStorage {
	return &InMemoryKeyValueStorage{
		booksMap: make(map[string]*models.Book),
	}
}

func (i *InMemoryKeyValueStorage) Insert(book *models.Book) (*models.Book, error) {
	i.booksMap[book.Isbn] = book
	return book, nil
}

func (i *InMemoryKeyValueStorage) GetAll() []*models.Book {
	books := make([]*models.Book, 0, len(i.booksMap))
	for _, value := range i.booksMap {
		books = append(books, value)
	}
	return books
}

func (i *InMemoryKeyValueStorage) SelectBookByIsbn(isbn string) (*models.Book, bool) {
	books, ok := i.booksMap[isbn]
	return books, ok
}

func (i *InMemoryKeyValueStorage) IncreaseBookViewsByIsbn(isbn string) {
	books := i.booksMap[isbn]
	books.Views++
	i.booksMap[isbn] = books
}
