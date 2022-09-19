package repositories

import "RSOI/src/catalogue/models"

type IBooksRepository interface {
	GetAll() ([]models.Book, error)
	UpdateBookLikes(uuid string, change int) error
	UpdateBookDislikes(uuid string, change int) error
}
