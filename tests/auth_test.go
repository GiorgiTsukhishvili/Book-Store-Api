package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
	"testing"

	"github.com/GiorgiTsukhishvili/BookShelf-Api/responses"
)

func LoginEndpointRequest(t *testing.T, email string, password string) responses.LoginResponse {
	reqBody := []byte(fmt.Sprintf(`{"email": "%s", "password": "%s"}`, email, password))

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
	response := LoginEndpointRequest(t, "admin@example.com", "admin123")

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
	loginResponse := LoginEndpointRequest(t, "admin@example.com", "admin123")

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

func TestRegisterEndpoint(t *testing.T) {
	reqBody := []byte(`{"name": "John Doe", "email": "johnDoe@example.com", "password": "password", "repeat_password": "password", "phone_number":"+995511111112", "type": "user"}`)

	req, err := http.NewRequest("POST", "http://localhost:3000/api/v1/register", bytes.NewBuffer(reqBody))
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

	var response responses.MessageResponse
	if err := json.NewDecoder(resp.Body).Decode(&response); err != nil {
		t.Fatalf("Could not decode response JSON: %v", err)
	}

	if response.Message == "" && response.Message != "Verification email was sent" {
		t.Errorf("Response does not contain expected keys or values: %+v", response)
	}
}

func TestLogoutEndpoint(t *testing.T) {
	loginResponse := LoginEndpointRequest(t, "admin@example.com", "admin123")

	req, err := http.NewRequest("POST", "http://localhost:3000/api/v1/logout", nil)
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

	if response.Message == "" && response.Message != "User logged out" {
		t.Errorf("Response does not contain expected keys or values: %+v", response)
	}
}
