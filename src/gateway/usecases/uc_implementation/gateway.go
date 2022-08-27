package uc_implementation

import (
	"RSOI/src/gateway/connector"
	"RSOI/src/gateway/gateway_error"
	"RSOI/src/gateway/models"
	"RSOI/src/gateway/usecases"
)

type GatewayUsecase struct {
	connector        connector.IGatewayConnector
	catalogueCB      usecases.CircuitBreaker
	usersCB          usecases.CircuitBreaker
	recommendationCB usecases.CircuitBreaker
}

func NewGatewayUsecase(connector connector.IGatewayConnector) GatewayUsecase {
	return GatewayUsecase{
		connector:        connector,
		catalogueCB:      *usecases.NewCircuitBreaker(50),
		usersCB:          *usecases.NewCircuitBreaker(50),
		recommendationCB: *usecases.NewCircuitBreaker(50),
	}
}

func (gc GatewayUsecase) GetRecommendations(lib []models.Book, prefs models.PreferencesList) ([]models.Book, gateway_error.GatewayError) {
	//TODO implement me
	panic("implement me")
}

func (gc GatewayUsecase) GetUserPreferences(uuid string) (models.PreferencesList, gateway_error.GatewayError) {
	code := gateway_error.Ok
	res, err := gc.usersCB.Call(func() (interface{}, error) { return gc.connector.GetUserPreferences(uuid) })
	if err != nil {
		code = gateway_error.Internal
	}
	return res.(models.PreferencesList), gateway_error.GatewayError{Err: err, Code: code}
}

func (gc GatewayUsecase) GetCatalogue() ([]models.Book, gateway_error.GatewayError) {
	//TODO implement me
	panic("implement me")
}
