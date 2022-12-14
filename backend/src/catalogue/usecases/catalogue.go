package usecases

import "RSOI/src/catalogue/models"

type IBooksUsecase interface {
	GetCatalogue() ([]models.Book, error)
	UpdateBookScore(uuid string, likesdiff, dislikesdiff int) error
}
