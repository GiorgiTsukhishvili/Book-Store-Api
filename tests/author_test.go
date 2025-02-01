package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"testing"

	"github.com/GiorgiTsukhishvili/BookShelf-Api/responses"
)

func TestAuthorGetEndpoint(t *testing.T) {
	loginResponse := LoginEndpointRequest(t)

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

	var response responses.AuthorGetResponse
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
