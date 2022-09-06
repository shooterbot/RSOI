package handlers

import (
	"RSOI/src/catalogue/usecases"
)

type BooksHandlers struct {
	bc usecases.IBooksUsecase
}

func NewBooksHandlers(booksCase usecases.IBooksUsecase) *BooksHandlers {
	return &BooksHandlers{bc: booksCase}
}
