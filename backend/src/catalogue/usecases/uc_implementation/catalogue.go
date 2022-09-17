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
