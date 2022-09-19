package repo_implementation

import (
	"RSOI/src/catalogue/models"
	"RSOI/src/database/pgdb"
	"RSOI/src/utility"
	"errors"
	"fmt"
)

const (
	selectAllBooks     = `select id, uid, name, publisher, year, likes, dislikes, status from books;`
	selectBookTags     = `select tag from books_tags where book_id=$1;`
	updateBookLikes    = `update books set likes = likes + $2 where uid = $1`
	updateBookDislikes = `update books set dislikes = dislikes + $2 where uid = $1`
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
			Status:    utility.BytesToBool(row[7]),
			Tags:      make([]string, 0),
		}
		likes := utility.BytesToInt(row[5])
		dislikes := utility.BytesToInt(row[6])
		if likes+dislikes == 0 {
			book.Rating = 50
		} else {
			book.Rating = 100 * likes / (likes + dislikes)
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

func (br *BooksRepository) UpdateBookLikes(uuid string, change int) error {
	affected, err := br.db.Exec(updateBookLikes, uuid, change)
	if err == nil {
		if affected == 0 {
			err = errors.New("Failed to update book likes: given UUID does not exist")
		}
	}
	return err
}

func (br *BooksRepository) UpdateBookDislikes(uuid string, change int) error {
	affected, err := br.db.Exec(updateBookDislikes, uuid, change)
	if err == nil {
		if affected == 0 {
			err = errors.New("Failed to update book dislikes: given UUID does not exist")
		}
	}
	return err
}
