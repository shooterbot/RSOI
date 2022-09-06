package usecases

import "RSOI/src/recommendations/models"

type IRecommendationsUsecase interface {
	GetRecommendations(lib []models.Book, prefs *models.PreferencesList) []models.Book
}
