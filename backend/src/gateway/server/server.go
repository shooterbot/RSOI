package server

import (
	"RSOI/src/gateway/connector"
	"RSOI/src/gateway/connector/connector_implementation"
	"RSOI/src/gateway/handlers"
	"RSOI/src/gateway/usecases/uc_implementation"
	"fmt"
	"github.com/gorilla/mux"
	"net/http"
)

func RunServer(address string, config *connector.Config) error {
	router := mux.NewRouter()
	apiRouter := router.PathPrefix("/api/v1").Subrouter()

	con := connector_implementation.NewGatewayConnector(config)
	gc := uc_implementation.NewGatewayUsecase(con)
	gh := handlers.NewGatewayHandlers(gc)
	//defer gc.Close()

	apiRouter.HandleFunc("/catalogue", gh.GetCatalogue).Methods(http.MethodGet)
	apiRouter.HandleFunc("/users", gh.CreateUser).Methods(http.MethodPost)
	apiRouter.HandleFunc("/sessions", gh.LoginUser).Methods(http.MethodPost)

	server := http.Server{
		Addr:    address,
		Handler: apiRouter,
	}

	fmt.Printf("Gateway service server is running on %s\n", address)
	return server.ListenAndServe()
}
