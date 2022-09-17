package repo_implementation

import (
	"RSOI/src/catalogue/models"
	"RSOI/src/database/pgdb"
	"RSOI/src/utility"
	"fmt"
)

const (
	selectAllBooks = `select id, uid, name, publisher, year, rating, status from books;`
	selectBookTags = `select tag from books_tags where book_id=$1;`
)

type BooksRepository struct {
	db *pgdb.DBManager
}

func NewBooksRepository(manager *pgdb.DBManager) *BooksRepository {
	return &BooksRepository{db: manager}
}

func (br *BooksRepository) GetAll() ([]models.Book, error) {
	data, err := br.db.Query(selectAllBooks)
	if err != nil {
		fmt.Printf("Failed to get books from db\n")
	}
	var res []models.Book

	for _, row := range data {
		book := models.Book{
			Id:        utility.BytesToInt(row[0]),
			Uuid:      utility.BytesToUid(row[1]),
			Name:      utility.BytesToString(row[2]),
			Publisher: utility.BytesToString(row[3]),
			Year:      utility.BytesToInt(row[4]),
			Rating:    utility.BytesToInt(row[5]),
			Status:    utility.BytesToBool(row[6]),
			Tags:      make([]string, 0),
		}
		tagsData, err := br.db.Query(selectBookTags, book.Id)
		if err != nil {
			fmt.Printf("Failed to get tags from db\n")
		}
		for _, tagsRow := range tagsData {
			book.Tags = append(book.Tags, utility.BytesToString(tagsRow[0]))
		}
		res = append(res, book)
	}
	return res, err
}
