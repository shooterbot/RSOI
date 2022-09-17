package server

import (
	"RSOI/src/database/pgdb"
	"RSOI/src/recommendations/handlers"
	"RSOI/src/recommendations/usecases/uc_implementation"
	"fmt"
	"github.com/gorilla/mux"
	"net/http"
	"os"
	"os/signal"
	"syscall"
)

func RunServer(address string, connectionString string) error {
	router := mux.NewRouter()
	apiRouter := router.PathPrefix("/api/v1").Subrouter()

	manager := pgdb.NewPGDBManager()
	err := manager.Connect(connectionString)
	if err != nil {
		fmt.Print("Failed to connect to database")
	} else {
		fmt.Println("Successfully connected to postgres database")
	}

	rh := handlers.NewRecommendationsHandlers(&uc_implementation.RecommendationsUsecase{})

	apiRouter.HandleFunc("/recommendations", rh.GetRecommendations).Methods(http.MethodGet)

	server := http.Server{
		Addr:    address,
		Handler: apiRouter,
	}

	sigs := make(chan os.Signal, 1)
	signal.Notify(sigs, syscall.SIGINT, syscall.SIGTERM)
	go func() {
		<-sigs
		manager.Disconnect()
		os.Exit(0)
	}()

	fmt.Printf("Recommendations system server is running on %s\n", address)
	return server.ListenAndServe()
}
