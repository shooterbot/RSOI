package connector_implementation

import (
	"RSOI/src/gateway/connector"
)

type GatewayConnector struct {
	config connector.Config
}

func NewGatewayConnector(config *connector.Config) *GatewayConnector {
	return &GatewayConnector{config: *config}
}
