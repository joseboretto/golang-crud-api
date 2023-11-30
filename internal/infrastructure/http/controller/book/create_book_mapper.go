package book

import "github.com/joseboretto/golang-crud-api/internal/domain/model"

func MapToDomain(c CreateBookRequest) *model.Book {
	return &model.Book{
		Title:      c.Title,
		TotalPages: c.TotalPages,
		Isbn:       c.Isbn,
	}
}
