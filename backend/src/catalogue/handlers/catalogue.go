package handlers

import (
	"RSOI/src/catalogue/usecases"
	"encoding/json"
	"fmt"
	"github.com/gorilla/mux"
	"io"
	"net/http"
	"strconv"
)

type BooksHandlers struct {
	bc usecases.IBooksUsecase
}

func NewBooksHandlers(booksCase usecases.IBooksUsecase) *BooksHandlers {
	return &BooksHandlers{bc: booksCase}
}

func writeError(w http.ResponseWriter, err string, status int) {
	w.WriteHeader(status)
	_, _ = w.Write([]byte(fmt.Sprintf(`{"message":"%s"}`, err)))
}

func (bh *BooksHandlers) GetCatalogue(w http.ResponseWriter, r *http.Request) {
	defer func(Body io.ReadCloser) {
		err := Body.Close()
		if err != nil {
			fmt.Println("Failed to close response body")
		}
	}(r.Body)

	rec, err := bh.bc.GetCatalogue()
	if err != nil {
		fmt.Println("Failed to get catalogue")
		writeError(w, "Error while getting catalogue", http.StatusServiceUnavailable)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	err = json.NewEncoder(w).Encode(rec)
	if err != nil {
		fmt.Println("Encoding json error: ", err)
		http.Error(w, "Failed to encode data to json", http.StatusInternalServerError)
		return
	}
}

func (bh *BooksHandlers) UpdateBookScore(w http.ResponseWriter, r *http.Request) {
	defer func(Body io.ReadCloser) {
		err := Body.Close()
		if err != nil {
			fmt.Println("Failed to close response body")
		}
	}(r.Body)

	var likesdiff, dislikesdiff int
	var err error

	vars := mux.Vars(r)
	bookUuid := vars["bookUuid"]
	if bookUuid == "" {
		fmt.Println("Received an invalid path parameter")
		writeError(w, "Failed to return the book: wrong path parameter", http.StatusBadRequest)
		return
	}
	likes := r.Header.Get("Likes")
	dislikes := r.Header.Get("Dislikes")
	likesdiff, err = strconv.Atoi(likes)
	if err == nil {
		dislikesdiff, err = strconv.Atoi(dislikes)
	}
	if err != nil {
		fmt.Println("Received an invalid header")
		writeError(w, "Failed to return the book: wrong header", http.StatusBadRequest)
	}
	err = bh.bc.UpdateBookScore(bookUuid, likesdiff, dislikesdiff)
	if err != nil {
		fmt.Println("Failed to update book score")
		writeError(w, "Failed to update book score", http.StatusInternalServerError)
	}
}
