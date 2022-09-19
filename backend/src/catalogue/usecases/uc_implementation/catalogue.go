package uc_implementation

import (
	"RSOI/src/catalogue/models"
	"RSOI/src/catalogue/repositories"
)

type BooksUsecase struct {
	br repositories.IBooksRepository
}

func NewBooksUsecase(repo repositories.IBooksRepository) *BooksUsecase {
	return &BooksUsecase{br: repo}
}

func (bc *BooksUsecase) GetCatalogue() ([]models.Book, error) {
	return bc.br.GetAll()
}

func (bc *BooksUsecase) UpdateBookScore(uuid string, likesdiff, dislikesdiff int) error {
	err := bc.br.UpdateBookLikes(uuid, likesdiff)
	if err == nil {
		err = bc.br.UpdateBookDislikes(uuid, dislikesdiff)
	}
	return err
}
