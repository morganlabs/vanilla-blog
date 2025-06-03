package main

import (
	"fmt"
	"log"
	"net/http"
	"os"
	"time"

	_ "github.com/joho/godotenv/autoload" // Used to autoload .env files
)

// Globally-accessible Variables
// All of these variables can be accessed throughout the entire `main` package
var startTime time.Time
var database Database

// Get the ADDRESS and PORT from environment
var address string = getEnv("ADDRESS", "127.0.0.1")
var port string = getEnv("PORT", "1234")

// The main "entry" function
func main() {
	// Record start time
	startTime = time.Now()

	startServer()
}

func startServer() {
	mux := initialiseMux()

	// Initialise the database
	db, err := newDatabase("database.sqlite")

	// Check for errors on start-up (I love this)
	if err != nil {
		fmt.Println("Failed to open database", err)
		os.Exit(-1)
	}

	// Assign the Database{} to the global `database` variable
	database = *db

	// Print a message about the server starting
	fmt.Printf("Starting server on %s:%s\n", address, port)

	// log.Fatal() - If the server fails, log it
	// ListenAndServe() - Starts the server on the specified address and port
	// Sprintf allows strings to be formatted with variables
	log.Fatal(http.ListenAndServe(fmt.Sprintf("%s:%s", address, port), mux))
}

func initialiseMux() *http.ServeMux {
	// Start a new HTTP Multiplexer
	// Google this idfk
	mux := http.NewServeMux()
	registerRoutes(mux)
	return mux
}
