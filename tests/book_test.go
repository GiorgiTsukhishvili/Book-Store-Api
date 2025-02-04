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

var BookId int

func TestBookGetEndpoint(t *testing.T) {
	loginResponse := LoginEndpointRequest(t, "business@example.com", "business123")

	req, err := http.NewRequest("GET", "http://localhost:3000/api/v1/book/1", nil)
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

	var response responses.BookRetrieveResponse
	if err := json.NewDecoder(resp.Body).Decode(&response); err != nil {
		t.Fatalf("Could not decode response JSON: %v", err)
	}

	if response.Book.Description == "" ||
		response.Book.Name == "" ||
		response.Book.Image == "" ||
		response.Book.Price == "" ||
		response.Book.CreationDate.IsZero() {
		t.Errorf("Response does not contain expected keys or values: %+v", response)
	}
}

func TestBooksGetEndpoint(t *testing.T) {
	loginResponse := LoginEndpointRequest(t, "business@example.com", "business123")

	req, err := http.NewRequest("GET", "http://localhost:3000/api/v1/book/?page=1&size=2", nil)
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

	var response responses.BooksGetResponse
	if err := json.NewDecoder(resp.Body).Decode(&response); err != nil {
		t.Fatalf("Could not decode response JSON: %v", err)
	}

	if response.Data == nil ||
		response.Pagination.CurrentPage == "" {
		t.Errorf("Response does not contain expected keys or values: %+v", response)
	}
}

func BookPostEndpoint(t *testing.T) {
	loginResponse := LoginEndpointRequest(t, "business@example.com", "business123")

	var requestBody bytes.Buffer
	writer := multipart.NewWriter(&requestBody)

	// Add form fields
	_ = writer.WriteField("name", "Jane Austen 1")
	_ = writer.WriteField("creation_date", "2006-01-02T15:04:05Z")
	_ = writer.WriteField("description", "English novelist known for her realism and social commentary.")
	_ = writer.WriteField("price", "9.99")
	_ = writer.WriteField("image_path", "pride_and_prejudice.jpg")
	_ = writer.WriteField("author_id", "1")
	genreIDs := []uint{1, 2, 3}
	for _, id := range genreIDs {
		_ = writer.WriteField("genre_ids", fmt.Sprintf("%d", id))
	}
	writer.Close()

	req, err := http.NewRequest("POST", "http://localhost:3000/api/v1/book/", &requestBody)
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

	var response responses.BookRetrieveResponse
	if err := json.NewDecoder(resp.Body).Decode(&response); err != nil {
		t.Fatalf("Could not decode response JSON: %v", err)
	}

	if response.Book.Description == "" ||
		response.Book.Name == "" ||
		response.Book.Image == "" ||
		response.Book.Price == "" ||
		response.Book.CreationDate.IsZero() {
		t.Errorf("Response does not contain expected keys or values: %+v", response)
	}

	BookId = int(response.Book.ID)
}

func BookPutEndpoint(t *testing.T) {
	loginResponse := LoginEndpointRequest(t, "business@example.com", "business123")

	var requestBody bytes.Buffer
	writer := multipart.NewWriter(&requestBody)

	// Add form fields
	_ = writer.WriteField("id", fmt.Sprint(BookId))
	_ = writer.WriteField("name", "Jane Austen 1")
	_ = writer.WriteField("creation_date", "2006-01-02T15:04:05Z")
	_ = writer.WriteField("description", "English novelist known for her realism and social commentary.")
	_ = writer.WriteField("price", "9.99")
	_ = writer.WriteField("image_path", "pride_and_prejudice.jpg")
	_ = writer.WriteField("author_id", "1")
	genreIDs := []uint{1, 2, 3}
	for _, id := range genreIDs {
		_ = writer.WriteField("genre_ids", fmt.Sprintf("%d", id))
	}
	writer.Close()

	req, err := http.NewRequest("PUT", "http://localhost:3000/api/v1/book/", &requestBody)
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

func BookDeleteEndpoint(t *testing.T) {
	loginResponse := LoginEndpointRequest(t, "business@example.com", "business123")

	req, err := http.NewRequest("DELETE", fmt.Sprintf("http://localhost:3000/api/v1/book/%v", BookId), nil)
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

func TestBookFlow(t *testing.T) {
	BookPostEndpoint(t)
	BookPutEndpoint(t)
	BookDeleteEndpoint(t)
}
