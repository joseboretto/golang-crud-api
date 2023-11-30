package dto

import "github.com/joseboretto/golang-crud-api/internal/domain/models"

func MapToGetAllBookResponse(m *models.Book) *GetAllBookResponse {
	return &GetAllBookResponse{
		Title:      m.Title,
		TotalPages: m.TotalPages,
		Isbn:       m.Isbn,
	}
}
