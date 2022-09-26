package handlers

import (
	"RSOI/src/gateway/gateway_error"
	"RSOI/src/gateway/models"
	"RSOI/src/gateway/usecases"
	"encoding/json"
	"fmt"
	"github.com/gorilla/mux"
	"io"
	"net/http"
)

type GatewayHandlers struct {
	gc usecases.IGatewayUsecase
}

func NewGatewayHandlers(gatewayCase usecases.IGatewayUsecase) *GatewayHandlers {
	return &GatewayHandlers{gc: gatewayCase}
}

func writeError(w http.ResponseWriter, err string, status int) {
	w.WriteHeader(status)
	_, _ = w.Write([]byte(fmt.Sprintf(`{"message":"%s"}`, err)))
}

func (gh *GatewayHandlers) GetRecommendations(w http.ResponseWriter, r *http.Request) {
	defer func(Body io.ReadCloser) {
		err := Body.Close()
		if err != nil {
			fmt.Println("Failed to close response body")
		}
	}(r.Body)

	query := r.URL.Query()
	userUuid, present := query["useruuid"]
	if !present || len(userUuid) != 1 {
		fmt.Println("Received an invalid query parameter")
		writeError(w, "Failed to get recommendations: wrong query parameter", http.StatusBadRequest)
		return
	}

	tokenUsername := r.Context().Value("Username")
	tokenUuid := r.Context().Value("UUID")
	if tokenUuid == "" || tokenUsername == "" || tokenUuid != userUuid[0] {
		fmt.Println("Failed to verify user authorization")
		http.Error(w, "User is not authorized", http.StatusUnauthorized)
		return

	}

	rec, ret := gh.gc.GetRecommendations(userUuid[0])
	if ret.Code != gateway_error.Ok {
		fmt.Println("Failed to get recommendations from internal service")
		writeError(w, "Error while getting recommendations", http.StatusServiceUnavailable)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	err := json.NewEncoder(w).Encode(rec)
	if err != nil {
		fmt.Println("Encoding json error: ", err)
		http.Error(w, "Failed to encode data to json", http.StatusInternalServerError)
		return
	}
}

func (gh *GatewayHandlers) GetCatalogue(w http.ResponseWriter, r *http.Request) {
	defer func(Body io.ReadCloser) {
		err := Body.Close()
		if err != nil {
			fmt.Println("Failed to close response body")
		}
	}(r.Body)

	rec, ret := gh.gc.GetCatalogue()
	if ret.Code != gateway_error.Ok {
		fmt.Println("Failed to get catalogue from internal service")
		writeError(w, "Error while getting catalogue", http.StatusServiceUnavailable)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	err := json.NewEncoder(w).Encode(rec)
	if err != nil {
		fmt.Println("Encoding json error: ", err)
		http.Error(w, "Failed to encode data to json", http.StatusInternalServerError)
		return
	}
}

func (gh *GatewayHandlers) CreateUser(w http.ResponseWriter, r *http.Request) {
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

	ret := gh.gc.CreateUser(user)

	if ret.Code == gateway_error.User || ret.Code == gateway_error.Internal {
		fmt.Println("Failed to create new user")
		if ret.Code == gateway_error.Internal {
			writeError(w, "Error while creating a user", http.StatusServiceUnavailable)
		} else {
			writeError(w, "Bad input given to create a user", http.StatusBadRequest)
		}
		return
	}
}

func (gh *GatewayHandlers) AddUserBookScore(w http.ResponseWriter, r *http.Request) {
	defer func(Body io.ReadCloser) {
		err := Body.Close()
		if err != nil {
			fmt.Println("Failed to close response body")
		}
	}(r.Body)

	vars := mux.Vars(r)
	bookUuid := vars["bookUuid"]
	if bookUuid == "" {
		fmt.Println("Received an invalid path parameter")
		writeError(w, "Failed to return the book: wrong path parameter", http.StatusBadRequest)
		return
	}

	username := r.Header.Get("User-Name")
	score := r.Header.Get("Score")

	tokenUsername := r.Context().Value("Username")
	tokenUuid := r.Context().Value("UUID")
	if tokenUuid == "" || tokenUsername == "" || tokenUsername != username {
		fmt.Println("Failed to verify user authorization")
		http.Error(w, "User is not authorized", http.StatusUnauthorized)
		return

	}

	ret := gh.gc.AddUserBookScore(bookUuid, username, score)

	if ret.Code == gateway_error.User || ret.Code == gateway_error.Internal {
		fmt.Println("Failed to get add user score to a book")
		if ret.Code == gateway_error.Internal {
			writeError(w, "Error while updating book user score", http.StatusServiceUnavailable)
		} else {
			writeError(w, "Bad input given to updating book user score", http.StatusBadRequest)
		}
		return
	}
}

type Session struct {
	UserToken string `json:"user-token"`
	Username  string `json:"username"`
	Id        int    `json:"id"`
}

func (gh *GatewayHandlers) LoginUser(w http.ResponseWriter, r *http.Request) {
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

	res, ret := gh.gc.LoginUser(user)

	if ret.Code == gateway_error.User || ret.Code == gateway_error.Internal {
		fmt.Println("Failed to authenficate a user")
		if ret.Code == gateway_error.Internal {
			writeError(w, "Error while authenficating a user", http.StatusServiceUnavailable)
		} else {
			writeError(w, "Bad input given to authenficate a user", http.StatusBadRequest)
		}
		return
	}
	if !res {
		fmt.Println("User failed authenfication")
		writeError(w, "Authenfication failed", http.StatusUnauthorized)
	}
	w.Header().Set("Content-Type", "application/json")
	session := &Session{
		UserToken: "1",
		Username:  "me",
		Id:        1,
	}
	err = json.NewEncoder(w).Encode(session)
	if err != nil {
		fmt.Println("Encoding json error: ", err)
		http.Error(w, "Failed to encode data to json", http.StatusInternalServerError)
		return
	}
}
