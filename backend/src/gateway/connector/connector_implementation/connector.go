package connector_implementation

import (
	"RSOI/src/gateway/connector"
	"RSOI/src/gateway/models"
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"net/http"
)

type GatewayConnector struct {
	config connector.Config
}

func NewGatewayConnector(config *connector.Config) *GatewayConnector {
	return &GatewayConnector{config: *config}
}

func (gc *GatewayConnector) GetUserPreferences(userUuid string) (*models.PreferencesList, error) {
	url := fmt.Sprintf(gc.config.UsersAddress+gc.config.ApiPath+"users/%s/preferences", userUuid)
	response, err := http.Get(url)
	if err != nil {
		fmt.Println("Failed to get preferences from internal service")
		err = errors.New("Failed to get preferences from internal service")
		return nil, err
	}

	defer func(Body io.ReadCloser) {
		err := Body.Close()
		if err != nil {
			fmt.Println("Failed to close response body")
		}
	}(response.Body)

	res := &models.PreferencesList{}
	err = json.NewDecoder(response.Body).Decode(res)
	if err != nil {
		fmt.Println("Failed to decode the json received from an internal service")
		err = errors.New("Internal decoding error")
		return nil, err
	}
	return res, nil
}

func (gc *GatewayConnector) GetCatalogue() (*[]models.Book, error) {
	url := fmt.Sprintf(gc.config.CatalogueAddress + gc.config.ApiPath + "catalogue")
	response, err := http.Get(url)
	if err != nil {
		fmt.Println("Failed to get comics from internal service")
		err = errors.New("Failed to get comics from internal service")
		return nil, err
	}

	defer func(Body io.ReadCloser) {
		err := Body.Close()
		if err != nil {
			fmt.Println("Failed to close response body")
		}
	}(response.Body)

	res := &[]models.Book{}
	err = json.NewDecoder(response.Body).Decode(res)
	if err != nil {
		fmt.Println("Failed to decode the json received from an internal service")
		err = errors.New("Internal decoding error")
		return nil, err
	}
	return res, nil
}

func (gc *GatewayConnector) GetRecommendations(books *[]models.Book, prefs *models.PreferencesList) (*[]models.Book, error) {
	url := fmt.Sprintf(gc.config.CatalogueAddress + gc.config.ApiPath)
	booksdata, err := json.Marshal(books)
	if err != nil {
		fmt.Println("Failed to encode input data")
		return nil, err
	}
	prefsdata, err := json.Marshal(prefs)
	if err != nil {
		fmt.Println("Failed to encode input data")
		return nil, err
	}
	request, err := http.NewRequest("GET", url, bytes.NewReader(append(booksdata, prefsdata...)))
	if err != nil {
		fmt.Println("Failed to create an http request")
		return nil, err
	}
	client := &http.Client{}
	response, err := client.Do(request)
	if err != nil {
		fmt.Println("Failed to get recommendations from internal service")
		return nil, err
	}

	var res *[]models.Book
	err = json.NewDecoder(response.Body).Decode(res)
	if err != nil {
		fmt.Println("Failed to decode data from internal service")
		err = errors.New("Decoding error")
		return nil, err
	}

	return res, err
}

func (gc *GatewayConnector) AddBookScore(bookUuid, score string) error {
	url := fmt.Sprintf(gc.config.CatalogueAddress+gc.config.ApiPath+"books/%s", bookUuid)
	request, err := http.NewRequest("PATCH", url, nil)
	if err != nil {
		fmt.Println("Failed to create an http request")
		return err
	}
	request.Header.Set("Score", score)

	client := &http.Client{}
	_, err = client.Do(request)
	if err != nil {
		fmt.Println("Internal service failed to update book rating")
	}
	return err
}

func (gc *GatewayConnector) AddUserScore(username string, bookUuid string, score string) error {
	url := fmt.Sprintf(gc.config.UsersAddress + gc.config.ApiPath + "preferences")
	request, err := http.NewRequest("PUT", url, nil)
	if err != nil {
		fmt.Println("Failed to create an http request")
		return err
	}
	request.Header.Set("Username", username)
	request.Header.Set("Book-UUID", bookUuid)
	request.Header.Set("Score", score)

	client := &http.Client{}
	_, err = client.Do(request)
	if err != nil {
		fmt.Println("Internal service failed to add user score")
	}
	return err
}

func (gc *GatewayConnector) CreateUser(user *models.User) error {
	url := fmt.Sprintf(gc.config.UsersAddress + gc.config.ApiPath + "users")
	data, err := json.Marshal(user)
	if err != nil {
		fmt.Println("Failed to encode input data")
		err = errors.New("Encoding error")
		return err
	}
	request, err := http.NewRequest("POST", url, bytes.NewReader(data))
	if err != nil {
		fmt.Println("Failed to create an http request")
		return err
	}

	client := &http.Client{}
	_, err = client.Do(request)
	if err != nil {
		fmt.Println("Internal service failed to add user score")
	}
	return err
}

func (gc *GatewayConnector) LoginUser(user *models.User) (bool, error) {
	url := fmt.Sprintf(gc.config.UsersAddress + gc.config.ApiPath + "sessions")
	data, err := json.Marshal(user)
	if err != nil {
		fmt.Println("Failed to encode input data")
		err = errors.New("Encoding error")
		return false, err
	}
	request, err := http.NewRequest("POST", url, bytes.NewReader(data))
	if err != nil {
		fmt.Println("Failed to create an http request")
		return false, err
	}

	client := &http.Client{}
	response, err := client.Do(request)
	if err != nil {
		fmt.Println("Internal service failed to add user score")
	}
	return response.StatusCode == http.StatusOK, err
}
