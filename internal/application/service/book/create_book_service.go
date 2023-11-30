package book

import (
	"github.com/joseboretto/golang-crud-api/internal/domain/model"
)

type CreateBookService struct {
	repository CreateBookRepositoryInterface
}

func NewCreateBookService(repository CreateBookRepositoryInterface) *CreateBookService {
	return &CreateBookService{
		repository: repository,
	}
}

func (s *CreateBookService) CreateBook(book *model.Book) (*model.Book, error) {
	// TODO: Add business logic here
	storedBook, err := s.repository.InsertBook(book)
	if err != nil {
		return nil, err
	}
	return storedBook, nil
}
