package connector

import "RSOI/src/gateway/models"

type IGatewayConnector interface {
	GetUserPreferences(userUuid string) (*models.PreferencesList, error)
	GetCatalogue() (*[]models.Book, error)
	GetRecommendations(books *[]models.Book, prefs *models.PreferencesList) (*[]models.Book, error)
}
