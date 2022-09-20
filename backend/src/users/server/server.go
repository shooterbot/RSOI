package server

import (
	"RSOI/src/database/pgdb"
	"RSOI/src/users/handlers"
	"RSOI/src/users/repositories/repo_implementation"
	"RSOI/src/users/usecases/uc_implementation"
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

	ur := repo_implementation.NewUsersRepository(manager)
	uc := uc_implementation.NewUsersUsecase(ur)
	uh := handlers.NewUsersHandlers(uc)

	apiRouter.HandleFunc("/users", uh.CreateUser).Methods(http.MethodPost)
	apiRouter.HandleFunc("/sessions", uh.LoginUser).Methods(http.MethodPost)
	apiRouter.HandleFunc("/preferences", uh.SetUserScore).Methods(http.MethodPut)
	apiRouter.HandleFunc("/users/{userUuid:[0-9|a-z|\\-]+}/preferences", uh.GetUserPreferences).Methods(http.MethodGet)

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

	fmt.Printf("Users system server is running on %s\n", address)
	return server.ListenAndServe()
}
