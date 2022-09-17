package repositories

import "RSOI/src/catalogue/models"

type IBooksRepository interface {
	GetAll() ([]models.Book, error)
}
