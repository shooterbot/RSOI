package main

import (
	"RSOI/src/gateway/connector"
	"RSOI/src/gateway/server"
	"errors"
	"fmt"
	"os"
)

func main() {
	var err error = nil
	var config connector.Config
	prod, exists := os.LookupEnv("PROD")
	var address string
	if exists && prod == "true" {
		config = connector.Config{
			CatalogueAddress:       "https://rsoi-catalogue-ivanov.herokuapp.com/",
			RecommendationsAddress: "https://rsoi-recommendations-ivanov.herokuapp.com/",
			UsersAddress:           "https://rsoi-users-ivanov.herokuapp.com/",
			ApiPath:                "api/v1/",
		}
		port, exists := os.LookupEnv("PORT")
		if exists {
			address = "0.0.0.0:" + port
		} else {
			fmt.Print("Failed to get port")
			err = errors.New("Failed to get port")
		}
	} else {
		address = "127.0.0.1:31337"
		config = connector.Config{
			CatalogueAddress:       "http://127.0.0.1:31338/",
			RecommendationsAddress: "http://127.0.0.1:31339/",
			UsersAddress:           "http://127.0.0.1:31340/",
			ApiPath:                "api/v1/",
		}
	}

	if err == nil {
		fmt.Printf("Starting server on %s\n", address)
		err = server.RunServer(address, &config)
		if err != nil {
			fmt.Print("Failed to start a server\n")
		}
	}
}
