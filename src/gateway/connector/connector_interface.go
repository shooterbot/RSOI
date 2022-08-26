package connector

import "RSOI/src/gateway/models"

type IGatewayConnector interface {
	GetCityLibraries(city string) (*[]models.Library, error)
}
