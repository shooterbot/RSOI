package usecases

import (
	"RSOI/src/gateway/gateway_error"
	"RSOI/src/gateway/models"
)

type IGatewayUsecase interface {
	GetRecommendations(userUuid string) (*[]models.Book, gateway_error.GatewayError)
	GetUserPreferences(uuid string) (*models.PreferencesList, gateway_error.GatewayError)
	GetCatalogue() (*[]models.Book, gateway_error.GatewayError)
}
