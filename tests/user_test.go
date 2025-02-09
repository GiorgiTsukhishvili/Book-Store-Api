package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"mime/multipart"
	"net/http"
	"testing"

	"github.com/GiorgiTsukhishvili/BookShelf-Api/responses"
)

func TestMeEndpoint(t *testing.T) {
	loginResponse := LoginEndpointRequest(t, "admin@example.com", "admin123")

	req, err := http.NewRequest("GET", "http://localhost:3000/api/v1/me", nil)
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

	var response responses.MeResponse
	if err := json.NewDecoder(resp.Body).Decode(&response); err != nil {
		t.Fatalf("Could not decode response JSON: %v", err)
	}

	if response.User.Email == "" ||
		response.User.PhoneNumber == "" ||
		response.User.ID == 0 ||
		response.User.Name == "" {
		t.Errorf("Response does not contain expected keys or values: %+v", response)
	}
}

func TestUserPasswordPostEndpoint(t *testing.T) {
	loginResponse := LoginEndpointRequest(t, "admin@example.com", "admin123")

	reqBody := []byte(`{"password": "admin123", "repeat_password": "admin123"}`)

	req, err := http.NewRequest("PUT", "http://localhost:3000/api/v1/user/password-update", bytes.NewBuffer(reqBody))
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

	if response.Message == "" && response.Message != "User password updated successfully" {
		t.Errorf("Response does not contain expected keys or values: %+v", response)
	}
}

func TestUserPutEndpoint(t *testing.T) {
	loginResponse := LoginEndpointRequest(t, "admin@example.com", "admin123")

	var requestBody bytes.Buffer
	writer := multipart.NewWriter(&requestBody)

	_ = writer.WriteField("name", "Admin")
	_ = writer.WriteField("image_path", "jane_austen.jpg")
	writer.Close()

	req, err := http.NewRequest("PUT", "http://localhost:3000/api/v1/user/", &requestBody)
	if err != nil {
		t.Fatalf("Could not create request: %v", err)
	}

	req.Header.Set("Content-Type", writer.FormDataContentType())
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

	if response.Message == "" && response.Message != "User updated successfully" {
		t.Errorf("Response does not contain expected keys or values: %+v", response)
	}
}

func TestUserEmailPostEndpoint(t *testing.T) {
	loginResponse := LoginEndpointRequest(t, "admin@example.com", "admin123")

	reqBody := []byte(`{"email": "admin@example.com"}`)

	req, err := http.NewRequest("POST", "http://localhost:3000/api/v1/user/update-email", bytes.NewBuffer(reqBody))
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

	if response.Message == "" && response.Message != "User email updated successfully" {
		t.Errorf("Response does not contain expected keys or values: %+v", response)
	}
}

func TestUserDeleteEndpoint(t *testing.T) {
	loginResponse := LoginEndpointRequest(t, "admin@example.com", "admin123")

	req, err := http.NewRequest("DELETE", "http://localhost:3000/api/v1/user/1", nil)
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

	if response.Message == "" && response.Message != "User deleted successfully" {
		t.Errorf("Response does not contain expected keys or values: %+v", response)
	}
}
