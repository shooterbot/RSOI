package connector

import "RSOI/src/gateway/models"

type IGatewayConnector interface {
	GetUserPreferences(userUuid string) (*models.PreferencesList, error)
	GetCatalogue() (*[]models.Book, error)
	GetRecommendations(books *[]models.Book, prefs *models.PreferencesList) (*[]models.Book, error)
	AddBookScore(bookUuid string, likesdiff, dislikesdiff int) error
	AddUserScore(username string, bookUuid string, score string) (bool, error)
	CreateUser(user *models.User) error
	LoginUser(user *models.User) (*models.Session, error)
}
