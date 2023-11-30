package dto

import "github.com/joseboretto/golang-crud-api/internal/domain/model"

func MapToGetAllBookResponse(m *model.Book) *GetAllBookResponse {
	return &GetAllBookResponse{
		Title:      m.Title,
		TotalPages: m.TotalPages,
		Isbn:       m.Isbn,
	}
}
