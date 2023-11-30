package book

import "github.com/joseboretto/golang-crud-api/internal/domain/model"

func MapToBookModel(c CreateBookRequest) *model.Book {
	return &model.Book{
		Title:      c.Title,
		TotalPages: c.TotalPages,
		Isbn:       c.Isbn,
	}
}

func MapToCreateBookResponse(m *model.Book) *CreateBookResponse {
	return &CreateBookResponse{
		Title:      m.Title,
		TotalPages: m.TotalPages,
		Isbn:       m.Isbn,
	}
}
