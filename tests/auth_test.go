package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
	"testing"

	"github.com/GiorgiTsukhishvili/BookShelf-Api/responses"
)

func LoginEndpointRequest(t *testing.T) responses.LoginResponse {
	reqBody := []byte(`{"email": "admin@example.com", "password": "admin123"}`)

	req, err := http.NewRequest("POST", "http://localhost:3000/api/v1/login", bytes.NewBuffer(reqBody))
	if err != nil {
		t.Fatalf("Could not create request: %v", err)
	}
	req.Header.Set("Content-Type", "application/json")

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		t.Fatalf("Could not send request: %v", err)
	}
	defer resp.Body.Close()

	var response responses.LoginResponse
	if err := json.NewDecoder(resp.Body).Decode(&response); err != nil {
		t.Fatalf("Could not decode response JSON: %v", err)
	}

	return response
}

func TestLoginEndpoint(t *testing.T) {
	response := LoginEndpointRequest(t)

	if response.JWT.Token == "" ||
		response.JWT.RefreshToken == "" ||
		response.JWT.TokenExpiration.IsZero() ||
		response.JWT.RefreshTokenExpiration.IsZero() ||
		response.User.Email == "" ||
		response.User.ID == 0 ||
		response.User.Name == "" {
		t.Errorf("Response does not contain expected keys or values: %+v", response)
	}
}

func TestRefreshTokenEndpoint(t *testing.T) {
	loginResponse := LoginEndpointRequest(t)

	reqBody := []byte(fmt.Sprintf(`{"refreshToken": "%s"}`, loginResponse.JWT.RefreshToken))

	req, err := http.NewRequest("POST", "http://localhost:3000/api/v1/refresh-token", bytes.NewBuffer(reqBody))
	if err != nil {
		t.Fatalf("Could not create request: %v", err)
	}
	req.Header.Set("Content-Type", "application/json")

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		t.Fatalf("Could not send request: %v", err)
	}
	defer resp.Body.Close()

	var response responses.RefreshTokenResponse
	if err := json.NewDecoder(resp.Body).Decode(&response); err != nil {
		t.Fatalf("Could not decode response JSON: %v", err)
	}

	if response.JWT.Token == "" ||
		response.JWT.RefreshToken == "" ||
		response.JWT.TokenExpiration.IsZero() ||
		response.JWT.RefreshTokenExpiration.IsZero() {
		t.Errorf("Response does not contain expected keys or values: %+v", response)
	}
}
