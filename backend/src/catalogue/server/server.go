package server

import (
	"RSOI/src/catalogue/handlers"
	"RSOI/src/catalogue/repositories/repo_implementation"
	"RSOI/src/catalogue/usecases/uc_implementation"
	"RSOI/src/database/pgdb"
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

	br := repo_implementation.NewBooksRepository(manager)
	bc := uc_implementation.NewBooksUsecase(br)
	bh := handlers.NewBooksHandlers(bc)

	apiRouter.HandleFunc("/catalogue", bh.GetCatalogue).Methods(http.MethodGet)

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

	fmt.Printf("Catalogue system server is running on %s\n", address)
	return server.ListenAndServe()
}
