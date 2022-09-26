package handlers

import (
	"RSOI/src/config"
	"RSOI/src/users/models"
	"RSOI/src/users/usecases"
	"encoding/json"
	"fmt"
	"github.com/gorilla/mux"
	"github.com/robbert229/jwt"
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

	user := &models.UserAuthData{}
	err := json.NewDecoder(r.Body).Decode(user)
	if err != nil {
		writeError(w, "Bad input given to user creation", http.StatusBadRequest)
	}

	err = uh.uc.CreateUser(user)

	if err != nil {
		fmt.Println("Failed to create new user")
		writeError(w, "Failed to create new user", http.StatusInternalServerError)
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

	userData := &models.UserAuthData{}
	err := json.NewDecoder(r.Body).Decode(userData)
	if err != nil {
		writeError(w, "Bad input given to user creation", http.StatusBadRequest)
	}

	user, err := uh.uc.LoginUser(userData)

	if err != nil {
		fmt.Println("Failed to authenficate the user")
		writeError(w, "Failed to authenficate the user", http.StatusInternalServerError)
		return
	}

	if user == nil {
		fmt.Println("UserAuthData failed authenfication")
		writeError(w, "Authenfication failed", http.StatusBadRequest)
		return
	}

	algorithm := jwt.HmacSha256(config.JWTKey)
	claims := jwt.NewClaim()
	claims.Set("UUID", user.UUID)
	claims.Set("Username", user.Username)

	token, err := algorithm.Encode(claims)
	if err != nil {
		fmt.Println(err)
		http.Error(w, `{"message":"cannot create jwt-token"}`, http.StatusInternalServerError)
		return
	}
	user.JWT = token

	w.Header().Set("Content-Type", "application/json")
	err = json.NewEncoder(w).Encode(user)
	if err != nil {
		fmt.Println("Encoding json error: ", err)
		http.Error(w, "Failed to encode data to json", http.StatusInternalServerError)
		return
	}
}

func (uh *UsersHandlers) GetUserPreferences(w http.ResponseWriter, r *http.Request) {
	defer func(Body io.ReadCloser) {
		err := Body.Close()
		if err != nil {
			fmt.Println("Failed to close response body")
		}
	}(r.Body)

	vars := mux.Vars(r)
	uUuid := vars["userUuid"]
	if uUuid == "" {
		fmt.Println("Received an invalid path parameter")
		http.Error(w, "Failed to get preferences: wrong path parameter", http.StatusBadRequest)
		return
	}

	prefs, err := uh.uc.GetUserPreferences(uUuid)

	if err != nil {
		fmt.Println("Failed to get user preferences")
		writeError(w, "Failed to get user preferences", http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	err = json.NewEncoder(w).Encode(prefs)
	if err != nil {
		fmt.Println("Encoding json error: ", err)
		http.Error(w, "Failed to encode data to json", http.StatusInternalServerError)
		return
	}
}

func (uh *UsersHandlers) SetUserScore(w http.ResponseWriter, r *http.Request) {
	defer func(Body io.ReadCloser) {
		err := Body.Close()
		if err != nil {
			fmt.Println("Failed to close response body")
		}
	}(r.Body)

	username := r.Header.Get("Username")
	bookUuid := r.Header.Get("Book-UUID")
	score := r.Header.Get("Score")

	changed, err := uh.uc.SetUserScore(username, bookUuid, score)

	if err != nil {
		fmt.Println("Failed to set user score")
		writeError(w, "Failed to set user score", http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	err = json.NewEncoder(w).Encode(changed)
	if err != nil {
		fmt.Println("Encoding json error: ", err)
		http.Error(w, "Failed to encode data to json", http.StatusInternalServerError)
		return
	}
}
