package main

import (
	"fmt"
	"io"
	"net/http"
)

// Function that retrieves SSL Labs API information
func getInfo(entryPoint string) (string, error) {
	url := entryPoint + "info"

	// Make HTTP GET request
	resp, err := http.Get(url)
	if err != nil {
		return "", err
	}

	// Ensure the response body is closed after function completes
	defer resp.Body.Close()

	// Check if the server didn't respond with a 200 OK status
	if resp.StatusCode != http.StatusOK {
		return "", fmt.Errorf("Server not available. Error: %s", resp.Status)
	}

	// Read the response body
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return "", err
	}

	return string(body), nil
}
