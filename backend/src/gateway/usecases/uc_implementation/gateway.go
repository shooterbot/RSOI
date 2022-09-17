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

func NewGatewayUsecase(connector connector.IGatewayConnector) *GatewayUsecase {
	return &GatewayUsecase{
		connector:        connector,
		catalogueCB:      *usecases.NewCircuitBreaker(50),
		usersCB:          *usecases.NewCircuitBreaker(50),
		recommendationCB: *usecases.NewCircuitBreaker(50),
	}
}

func (gc *GatewayUsecase) GetRecommendations(userUuid string) (*[]models.Book, gateway_error.GatewayError) {
	var res *[]models.Book = nil
	prefs, gerr := gc.GetUserPreferences(userUuid)
	if gerr.Code != gateway_error.Ok {
		var books *[]models.Book
		books, gerr = gc.GetCatalogue()
		if gerr.Code != gateway_error.Ok {
			rec, err := gc.recommendationCB.Call(func() (interface{}, error) { return gc.connector.GetRecommendations(books, prefs) })
			if err == nil {
				res = rec.(*[]models.Book)
			} else {
				gerr = gateway_error.GatewayError{Err: err, Code: gateway_error.Internal}
			}
		}
	}
	return res, gerr
}

func (gc *GatewayUsecase) GetUserPreferences(uuid string) (*models.PreferencesList, gateway_error.GatewayError) {
	var res *models.PreferencesList
	code := gateway_error.Ok
	prefs, err := gc.usersCB.Call(func() (interface{}, error) { return gc.connector.GetUserPreferences(uuid) })
	if err != nil {
		res = nil
		code = gateway_error.Internal
	} else {
		res = prefs.(*models.PreferencesList)
	}
	return res, gateway_error.GatewayError{Err: err, Code: code}
}

func (gc *GatewayUsecase) GetCatalogue() (*[]models.Book, gateway_error.GatewayError) {
	var res *[]models.Book
	code := gateway_error.Ok
	books, err := gc.catalogueCB.Call(func() (interface{}, error) { return gc.connector.GetCatalogue() })
	if err != nil {
		code = gateway_error.Internal
	} else {
		res = books.(*[]models.Book)
	}
	return res, gateway_error.GatewayError{Err: err, Code: code}
}

func (gc *GatewayUsecase) AddUserBookScore(bookUuid string, username string, score string) gateway_error.GatewayError {
	code := gateway_error.Ok
	err := gc.connector.AddBookScore(bookUuid, score)
	if err != nil {
		code = gateway_error.Internal
	} else {
		err = gc.connector.AddUserScore(username, bookUuid, score)
		if err != nil {
			code = gateway_error.Internal
		}
	}
	return gateway_error.GatewayError{Err: err, Code: code}
}

func (gc *GatewayUsecase) CreateUser(user *models.User) gateway_error.GatewayError {
	code := gateway_error.Ok
	err := gc.connector.CreateUser(user)
	if err != nil {
		code = gateway_error.Internal
	}
	return gateway_error.GatewayError{Err: err, Code: code}
}
