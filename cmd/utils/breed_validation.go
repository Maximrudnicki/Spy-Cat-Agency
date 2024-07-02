package utils

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
)

type Breed struct {
	ID          string `json:"id"`
    Name        string `json:"name"`
    Description string `json:"description"`
}

func ValidateBreed(breed string) bool {
	url := "https://api.thecatapi.com/v1/breeds"

	req, err := http.NewRequest(http.MethodGet, url, nil)
	if err != nil {
		fmt.Printf("client: could not create request: %s\n", err)
		return false
	}

	res, err := http.DefaultClient.Do(req)
	if err != nil {
		fmt.Printf("client: error making http request: %s\n", err)
		return false
	}

	resBody, err := io.ReadAll(res.Body)
	if err != nil {
		fmt.Printf("client: could not read response body: %s\n", err)
		return false
	}

	var breeds []Breed
    err = json.Unmarshal(resBody, &breeds)
    if err != nil {
		fmt.Printf("client: could not unmarshal json: %s\n", err)
        return false
    }

    // A map to store breed names for efficient lookup
    allBreeds := make(map[string]bool)
    for _, breed := range breeds {
        allBreeds[breed.Name] = true
    }

	// Check if breed exists
    _, exists := allBreeds[breed]

	return exists
}
