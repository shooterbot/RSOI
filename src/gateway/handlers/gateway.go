package handlers

import (
	"RSOI/src/gateway/usecases"
	"encoding/json"
	"fmt"
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

	prefs := gh.gc.GetUserPreferences(userUuid[0])
	books := gh.gc.GetCatalogue()

	rec := gh.gc.GetRecommendations(books, prefs)

	w.Header().Set("Content-Type", "application/json")
	err := json.NewEncoder(w).Encode(rec)
	if err != nil {
		fmt.Println("Encoding json error: ", err)
		http.Error(w, "Failed to encode data to json", http.StatusInternalServerError)
		return
	}
}
