package uc_implementation

import (
	"RSOI/src/gateway/connector"
	"RSOI/src/gateway/models"
)

type GatewayUsecase struct {
	gc connector.IGatewayConnector
}

func NewGatewayUsecase(gc connector.IGatewayConnector) GatewayUsecase {
	return GatewayUsecase{gc: gc}
}

func (gc GatewayUsecase) GetRecommendations(lib []models.Book, prefs models.PreferencesList) []models.Book {
	//TODO implement me
	panic("implement me")
}

func (gc GatewayUsecase) GetUserPreferences(uuid string) models.PreferencesList {
	//TODO implement me
	panic("implement me")
}

func (gc GatewayUsecase) GetCatalogue() []models.Book {
	//TODO implement me
	panic("implement me")
}
