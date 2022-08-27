package connector

import "RSOI/src/gateway/models"

type IGatewayConnector interface {
	GetUserPreferences(userUuid string) (*models.PreferencesList, error)
}
