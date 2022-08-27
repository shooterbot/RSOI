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

func (gc GatewayUsecase) GetRecommendations(userUuid string) (*[]models.Book, gateway_error.GatewayError) {
	var res *[]models.Book = nil
	prefs, gerr := gc.GetUserPreferences(userUuid)
	if gerr.Code != gateway_error.Ok {
		var books *[]models.Book
		books, gerr = gc.GetCatalogue()
		if gerr.Code != gateway_error.Ok {
			rec, err := gc.recommendationCB.Call(gc.connector.GetRecommendations(books, prefs))
			if err == nil {
				res = rec.(*[]models.Book)
			}
		}
	}
	return res, gerr
}

func (gc GatewayUsecase) GetUserPreferences(uuid string) (*models.PreferencesList, gateway_error.GatewayError) {
	code := gateway_error.Ok
	var res *models.PreferencesList
	prefs, err := gc.usersCB.Call(func() (interface{}, error) { return gc.connector.GetUserPreferences(uuid) })
	if err != nil {
		res = nil
		code = gateway_error.Internal
	} else {
		res = prefs.(*models.PreferencesList)
	}
	return res, gateway_error.GatewayError{Err: err, Code: code}
}

func (gc GatewayUsecase) GetCatalogue() (*[]models.Book, gateway_error.GatewayError) {
	//TODO implement me
	panic("implement me")
}
