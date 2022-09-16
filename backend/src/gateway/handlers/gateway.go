package handlers

import (
	"RSOI/src/gateway/gateway_error"
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