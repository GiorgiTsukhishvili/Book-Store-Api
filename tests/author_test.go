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

var AuthorId int

func TestAuthorGetEndpoint(t *testing.T) {
	loginResponse := LoginEndpointRequest(t, "admin@example.com", "admin123")

	req, err := http.NewRequest("GET", "http://localhost:3000/api/v1/author/1", nil)
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

	var response responses.AuthorRetrieveResponse
	if err := json.NewDecoder(resp.Body).Decode(&response); err != nil {
		t.Fatalf("Could not decode response JSON: %v", err)
	}

	if response.Author.Description == "" ||
		response.Author.Name == "" ||
		response.Author.Image == "" ||
		response.Author.Nationality == "" ||
		response.Author.BirthDate.IsZero() {
		t.Errorf("Response does not contain expected keys or values: %+v", response)
	}
}

func TestAuthorsGetEndpoint(t *testing.T) {
	loginResponse := LoginEndpointRequest(t, "admin@example.com", "admin123")

	req, err := http.NewRequest("GET", "http://localhost:3000/api/v1/author/?page=1&size=2", nil)
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

	var response responses.AuthorsGetResponse
	if err := json.NewDecoder(resp.Body).Decode(&response); err != nil {
		t.Fatalf("Could not decode response JSON: %v", err)
	}

	if response.Data == nil ||
		response.Pagination.CurrentPage == "" {
		t.Errorf("Response does not contain expected keys or values: %+v", response)
	}
}

func AuthorPostEndpoint(t *testing.T) {
	loginResponse := LoginEndpointRequest(t, "admin@example.com", "admin123")

	var requestBody bytes.Buffer
	writer := multipart.NewWriter(&requestBody)

	// Add form fields
	_ = writer.WriteField("name", "Jane Austen 1")
	_ = writer.WriteField("birth_date", "2006-01-02T15:04:05Z")
	_ = writer.WriteField("description", "English novelist known for her realism and social commentary.")
	_ = writer.WriteField("nationality", "British")
	_ = writer.WriteField("image_path", "jane_austen.jpg")
	writer.Close()

	req, err := http.NewRequest("POST", "http://localhost:3000/api/v1/author/", &requestBody)
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

	var response responses.AuthorRetrieveResponse
	if err := json.NewDecoder(resp.Body).Decode(&response); err != nil {
		t.Fatalf("Could not decode response JSON: %v", err)
	}

	if response.Author.Description == "" ||
		response.Author.Name == "" ||
		response.Author.Image == "" ||
		response.Author.Nationality == "" ||
		response.Author.BirthDate.IsZero() {
		t.Errorf("Response does not contain expected keys or values: %+v", response)
	}

	AuthorId = int(response.Author.ID)
}

func AuthorPutEndpoint(t *testing.T) {
	loginResponse := LoginEndpointRequest(t, "admin@example.com", "admin123")

	var requestBody bytes.Buffer
	writer := multipart.NewWriter(&requestBody)

	_ = writer.WriteField("id", fmt.Sprint(AuthorId))
	_ = writer.WriteField("name", "Jane Austen 2")
	_ = writer.WriteField("birth_date", "2006-01-02T15:04:05Z")
	_ = writer.WriteField("description", "English novelist known for her realism and social commentary.")
	_ = writer.WriteField("nationality", "British")
	_ = writer.WriteField("image_path", "jane_austen.jpg")
	writer.Close()

	req, err := http.NewRequest("PUT", "http://localhost:3000/api/v1/author/", &requestBody)
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

	if response.Message == "" && response.Message != "Author updated successfully" {
		t.Errorf("Response does not contain expected keys or values: %+v", response)
	}

}

func AuthorDeleteEndpoint(t *testing.T) {
	loginResponse := LoginEndpointRequest(t, "admin@example.com", "admin123")

	req, err := http.NewRequest("DELETE", fmt.Sprintf("http://localhost:3000/api/v1/author/%v", AuthorId), nil)
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

	if response.Message == "" && response.Message != "Author deleted successfully" {
		t.Errorf("Response does not contain expected keys or values: %+v", response)
	}
}

func TestAuthorFlow(t *testing.T) {
	AuthorPostEndpoint(t)
	AuthorPutEndpoint(t)
	AuthorDeleteEndpoint(t)
}
