package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"os"
	"time"
)

func writeJSON(w http.ResponseWriter, status int, response any) {
	// Set the header type to application/json, to indicate a JSON response
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(status)

	// `err :=` means that this function *may* return an error
	// writeJSON() can be found below
	err := json.NewEncoder(w).Encode(response)

	// If `err` is not `nil`, an error has occured
	// Handle the error
	if err != nil {
		writeError(w, 500, "Failed to encode response")
	}
}

func writeError(w http.ResponseWriter, status int, message string) {
	// Create an instance of ApiError
	apiError := ApiError{status, message}

	// Use the writeJSON function to write the error to the screen
	writeJSON(w, status, apiError)
}

// A helper function to get environment variables with a fallback
func getEnv(key string, fallback string) string {
	// Attempt to get the env variable
	envVar := os.Getenv(key)

	// If it's empty, return the fallback value
	if envVar == "" {
		return fallback
	}
	return envVar
}

// A helper to format time.Duration into DD:HH:MM:SS
func formatDuration(d time.Duration) string {
	days := int(d.Hours()) / 24
	hours := int(d.Hours()) % 24
	minutes := int(d.Minutes()) % 60
	seconds := int(d.Seconds()) % 60
	return fmt.Sprintf("%02d:%02d:%02d:%02d", days, hours, minutes, seconds)
}
