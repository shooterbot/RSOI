package handlers

import (
	"RSOI/src/users/models"
	"RSOI/src/users/usecases"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
)

type UsersHandlers struct {
	uc usecases.IUsersUsecase
}

func NewUsersHandlers(usersCase usecases.IUsersUsecase) *UsersHandlers {
	return &UsersHandlers{uc: usersCase}
}

func writeError(w http.ResponseWriter, err string, status int) {
	w.WriteHeader(status)
	_, _ = w.Write([]byte(fmt.Sprintf(`{"message":"%s"}`, err)))
}

func (uh *UsersHandlers) CreateUser(w http.ResponseWriter, r *http.Request) {
	defer func(Body io.ReadCloser) {
		err := Body.Close()
		if err != nil {
			fmt.Println("Failed to close response body")
		}
	}(r.Body)

	user := &models.User{}
	err := json.NewDecoder(r.Body).Decode(user)
	if err != nil {
		writeError(w, "Bad input given to user creation", http.StatusBadRequest)
	}

	err = uh.uc.CreateUser(user)

	if err != nil {
		fmt.Println("Failed to create new user")
		writeError(w, "Bad input given to create a user", http.StatusInternalServerError)
		return
	}
}

func (uh *UsersHandlers) LoginUser(w http.ResponseWriter, r *http.Request) {
	defer func(Body io.ReadCloser) {
		err := Body.Close()
		if err != nil {
			fmt.Println("Failed to close response body")
		}
	}(r.Body)

	user := &models.User{}
	err := json.NewDecoder(r.Body).Decode(user)
	if err != nil {
		writeError(w, "Bad input given to user creation", http.StatusBadRequest)
	}

	checked, err := uh.uc.LoginUser(user)

	if err != nil {
		fmt.Println("Failed to create new user")
		writeError(w, "Bad input given to create a user", http.StatusInternalServerError)
		return
	}

	if !checked {
		fmt.Println("User failed authenfication")
		writeError(w, "Authenfication failed", http.StatusBadRequest)
	}
}
