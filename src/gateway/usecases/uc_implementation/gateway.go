package uc_implementation

import (
	"RSOI/src/gateway/connector"
	"RSOI/src/gateway/models"
	"RSOI/src/gateway/usecases"
)

type GatewayUsecase struct {
	connector     connector.IGatewayConnector
	libraryCB     usecases.CircuitBreaker
	reservationCB usecases.CircuitBreaker
	ratingCB      usecases.CircuitBreaker
}

func NewGatewayUsecase(connector connector.IGatewayConnector) GatewayUsecase {
	return GatewayUsecase{
		connector:     connector,
		libraryCB:     *usecases.NewCircuitBreaker(50),
		reservationCB: *usecases.NewCircuitBreaker(50),
		ratingCB:      *usecases.NewCircuitBreaker(50),
	}
}

func (gc GatewayUsecase) GetRecommendations(lib []models.Book, prefs models.PreferencesList) []models.Book {
	//TODO implement me
	panic("implement me")
}

func (gc GatewayUsecase) GetUserPreferences(uuid string) models.PreferencesList {

	panic("implement me")
}

func (gc GatewayUsecase) GetCatalogue() []models.Book {
	//TODO implement me
	panic("implement me")
}
