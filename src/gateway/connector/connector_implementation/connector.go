package connector_implementation

import (
	"RSOI/src/gateway/connector"
	"RSOI/src/gateway/models"
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"net/http"
)

type GatewayConnector struct {
	config connector.Config
}

func NewGatewayConnector(config *connector.Config) *GatewayConnector {
	return &GatewayConnector{config: *config}
}

func (gc *GatewayConnector) GetUserPreferences(userUuid string) (*models.PreferencesList, error) {
	url := fmt.Sprintf(gc.config.UsersAddress+gc.config.ApiPath+"users/%s/preferences", userUuid)
	response, err := http.Get(url)
	if err != nil {
		fmt.Println("Failed to get preferences from internal service")
		err = errors.New("Failed to get preferences from internal service")
		return nil, err
	}

	defer func(Body io.ReadCloser) {
		err := Body.Close()
		if err != nil {
			fmt.Println("Failed to close response body")
		}
	}(response.Body)

	res := &models.PreferencesList{}
	err = json.NewDecoder(response.Body).Decode(res)
	if err != nil {
		fmt.Println("Failed to decode the json received from an internal service")
		err = errors.New("Internal decoding error")
		return nil, err
	}
	return res, nil
}
