package usecases

import "RSOI/src/recommendations/models"

type IRecommendationsUsecase interface {
	GetRecommendations(city string) *models.PreferencesList
}
