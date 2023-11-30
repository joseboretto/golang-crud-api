package book

import (
	"github.com/joseboretto/golang-crud-api/internal/domain/model"
)

type GetAllBooksService struct {
	repository GetAllBooksRepositoryInterface
}

func NewGetAllBooksService(repository GetAllBooksRepositoryInterface) *GetAllBooksService {
	return &GetAllBooksService{
		repository: repository,
	}
}

func (s *GetAllBooksService) GetAllBooks() ([]*model.Book, error) {
	books, err := s.repository.SelectAllBooks()
	if err != nil {
		return nil, err
	}
	return books, nil
}
