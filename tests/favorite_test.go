package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
	"testing"

	"github.com/GiorgiTsukhishvili/BookShelf-Api/responses"
)

var FavoriteId int

func TestFavoritesGetEndpoint(t *testing.T) {
	loginResponse := LoginEndpointRequest(t)

	req, err := http.NewRequest("GET", "http://localhost:3000/api/v1/favorite/?page=1&size=2", nil)
	if err != nil {
		t.Fatalf("Could not create request: %v", err)
	}
	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("Authorization", fmt.Sprintf("Bearer %s", loginResponse.JWT.Token))

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		t.Fatalf("Could not send request: %v", err)
	}
	defer resp.Body.Close()

	var response responses.FavoritesGetResponse
	if err := json.NewDecoder(resp.Body).Decode(&response); err != nil {
		t.Fatalf("Could not decode response JSON: %v", err)
	}

	if response.Data == nil ||
		response.Pagination.CurrentPage == "" {
		t.Errorf("Response does not contain expected keys or values: %+v", response)
	}
}

func FavoritePostEndpoint(t *testing.T) {
	loginResponse := LoginEndpointRequest(t)

	reqBody := []byte(`{"book_id": "1"}`)

	req, err := http.NewRequest("POST", "http://localhost:3000/api/v1/favorite/", bytes.NewBuffer(reqBody))
	if err != nil {
		t.Fatalf("Could not create request: %v", err)
	}
	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("Authorization", fmt.Sprintf("Bearer %s", loginResponse.JWT.Token))

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		t.Fatalf("Could not send request: %v", err)
	}
	defer resp.Body.Close()

	var response responses.FavoritePostResponse
	if err := json.NewDecoder(resp.Body).Decode(&response); err != nil {
		t.Fatalf("Could not decode response JSON: %v", err)
	}

	if response.Favorite.BookID == 0 ||
		response.Favorite.UserID == 0 {
		t.Errorf("Response does not contain expected keys or values: %+v", response)
	}

	FavoriteId = int(response.Favorite.ID)
}

func FavoriteDeleteEndpoint(t *testing.T) {
	loginResponse := LoginEndpointRequest(t)

	req, err := http.NewRequest("DELETE", fmt.Sprintf("http://localhost:3000/api/v1/favorite/%v", FavoriteId), nil)
	if err != nil {
		t.Fatalf("Could not create request: %v", err)
	}
	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("Authorization", fmt.Sprintf("Bearer %s", loginResponse.JWT.Token))

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		t.Fatalf("Could not send request: %v", err)
	}
	defer resp.Body.Close()

	var response responses.MessageResponse
	if err := json.NewDecoder(resp.Body).Decode(&response); err != nil {
		t.Fatalf("Could not decode response JSON: %v", err)
	}

	if response.Message == "" && response.Message != "Favorite deleted successfully" {
		t.Errorf("Response does not contain expected keys or values: %+v", response)
	}
}

func TestFavoriteFlow(t *testing.T) {
	FavoritePostEndpoint(t)
	FavoriteDeleteEndpoint(t)
}
