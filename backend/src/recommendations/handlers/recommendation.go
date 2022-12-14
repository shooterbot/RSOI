package handlers

import (
	"RSOI/src/recommendations/models"
	"RSOI/src/recommendations/usecases"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
)

type RecommendationsHandlers struct {
	rc usecases.IRecommendationsUsecase
}

func NewRecommendationsHandlers(recommendationsCase usecases.IRecommendationsUsecase) *RecommendationsHandlers {
	return &RecommendationsHandlers{rc: recommendationsCase}
}

func writeError(w http.ResponseWriter, err string, status int) {
	w.WriteHeader(status)
	_, _ = w.Write([]byte(fmt.Sprintf(`{"message":"%s"}`, err)))
}

func (rh *RecommendationsHandlers) GetRecommendations(w http.ResponseWriter, r *http.Request) {
	defer func(Body io.ReadCloser) {
		err := Body.Close()
		if err != nil {
			fmt.Println("Failed to close response body")
		}
	}(r.Body)

	decoder := json.NewDecoder(r.Body)
	info := &models.RecomendationsInfo{}
	err := decoder.Decode(info)
	if err != nil {
		fmt.Println("Failed to decode the received json")
		writeError(w, "Bad json given as input", http.StatusBadRequest)
		return
	}

	rec := rh.rc.GetRecommendations(info.Books, &info.Prefs)

	w.Header().Set("Content-Type", "application/json")
	err = json.NewEncoder(w).Encode(rec)
	if err != nil {
		fmt.Println("Encoding json error: ", err)
		http.Error(w, "Failed to encode data to json", http.StatusInternalServerError)
		return
	}
}
