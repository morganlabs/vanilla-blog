package main

import (
	"fmt"
	_ "github.com/joho/godotenv/autoload" // Used to autoload .env files
	"log"
	"net/http"
	"time"
)

// Globally-accessible Variables
// All of these variables can be accessed throughout the entire `main` package
var startTime time.Time

// Get the ADDRESS and PORT from environment
var address string = getEnv("ADDRESS", "127.0.0.1")
var port string = getEnv("PORT", "1234")

// The main "entry" function
func main() {
	// Adds a handler when navigating to /api
	startServer()
}

func startServer() {

	// Record start time
	startTime = time.Now()

	// Start a new HTTP Multiplexer
	// Google this idfk
	mux := http.NewServeMux()
	registerRoutes(mux)

	// Print a message about the server starting
	fmt.Printf("Starting server on %s:%s\n", address, port)

	// log.Fatal() - If the server fails, log it
	// ListenAndServe() - Starts the server on the specified address and port
	// Sprintf allows strings to be formatted with variables
	log.Fatal(http.ListenAndServe(fmt.Sprintf("%s:%s", address, port), mux))
}
