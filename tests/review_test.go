package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
	"testing"

	"github.com/GiorgiTsukhishvili/BookShelf-Api/responses"
)

var ReviewId int

func TestReviewGetEndpoint(t *testing.T) {
	loginResponse := LoginEndpointRequest(t)

	req, err := http.NewRequest("GET", "http://localhost:3000/api/v1/review/1", nil)
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

	var response responses.ReviewRetrieveResponse
	if err := json.NewDecoder(resp.Body).Decode(&response); err != nil {
		t.Fatalf("Could not decode response JSON: %v", err)
	}

	if response.Review.Rating == "" ||
		response.Review.Comment == "" {
		t.Errorf("Response does not contain expected keys or values: %+v", response)
	}
}

func TestReviewsGetEndpoint(t *testing.T) {
	loginResponse := LoginEndpointRequest(t)

	req, err := http.NewRequest("GET", "http://localhost:3000/api/v1/review/?page=1&size=2", nil)
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

	var response responses.ReviewsGetResponse
	if err := json.NewDecoder(resp.Body).Decode(&response); err != nil {
		t.Fatalf("Could not decode response JSON: %v", err)
	}

	if response.Data == nil ||
		response.Pagination.CurrentPage == "" {
		t.Errorf("Response does not contain expected keys or values: %+v", response)
	}
}

func ReviewPostEndpoint(t *testing.T) {
	loginResponse := LoginEndpointRequest(t)

	reqBody := []byte(`{"rating": "4", "comment": "nice", "book_id": "1"}`)

	req, err := http.NewRequest("POST", "http://localhost:3000/api/v1/review/", bytes.NewBuffer(reqBody))
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

	var response responses.ReviewRetrieveResponse
	if err := json.NewDecoder(resp.Body).Decode(&response); err != nil {
		t.Fatalf("Could not decode response JSON: %v", err)
	}

	if response.Review.Rating == "" ||
		response.Review.Comment == "" {
		t.Errorf("Response does not contain expected keys or values: %+v", response)
	}

	ReviewId = int(response.Review.ID)
}

func ReviewPutEndpoint(t *testing.T) {
	loginResponse := LoginEndpointRequest(t)

	reqBody := []byte(fmt.Sprintf(`{"rating": "4", "comment": "nice", "book_id": "1", "id": "%v"}`, ReviewId))

	req, err := http.NewRequest("PUT", "http://localhost:3000/api/v1/review/", bytes.NewBuffer(reqBody))
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

	if response.Message == "" && response.Message != "Review updated successfully" {
		t.Errorf("Response does not contain expected keys or values: %+v", response)
	}
}

func TestReviewFlow(t *testing.T) {
	ReviewPostEndpoint(t)
	ReviewPutEndpoint(t)
}
