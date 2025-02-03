package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
	"testing"

	"github.com/GiorgiTsukhishvili/BookShelf-Api/responses"
)

func TestNotificationsGetEndpoint(t *testing.T) {
	loginResponse := LoginEndpointRequest(t)

	req, err := http.NewRequest("GET", "http://localhost:3000/api/v1/notification/?page=1&size=2", nil)
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

	var response responses.NotificationsGetResponse
	if err := json.NewDecoder(resp.Body).Decode(&response); err != nil {
		t.Fatalf("Could not decode response JSON: %v", err)
	}

	if response.Data == nil ||
		response.Pagination.CurrentPage == "" {
		t.Errorf("Response does not contain expected keys or values: %+v", response)
	}
}

func TestNotificationPutEndpoint(t *testing.T) {
	loginResponse := LoginEndpointRequest(t)

	reqBody := []byte(`{"id": ["1"]}`)

	req, err := http.NewRequest("PUT", "http://localhost:3000/api/v1/notification/", bytes.NewBuffer(reqBody))
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

	if response.Message == "" && response.Message != "Notification updated successfully" {
		t.Errorf("Response does not contain expected keys or values: %+v", response)
	}
}
