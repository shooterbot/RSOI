package usecases

import "RSOI/src/gateway/models"

type IGatewayUsecase interface {
	GetRecommendations(lib []models.Book, prefs models.PreferencesList) []models.Book
	GetUserPreferences(uuid string) models.PreferencesList
	GetCatalogue() []models.Book
}
