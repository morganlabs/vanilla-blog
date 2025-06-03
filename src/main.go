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

	// Initialise the database
	db, err := newDatabase("database.sqlite")

	if err != nil {
		fmt.Println("Failed to open database", err)
		os.Exit(-1)
	}

	database = *db

	datetime := time.Now().Format(time.ANSIC)
	_, err = database.NewPost("Title", "Summary", `
# Lens Types
		
* Single Vision (Distance/Intermediate/Near)
  Used to help with one aspect of vision, such as [[Long-sightedness (Hypermetropia)|Hypermetropia]] or [[Short-sightedness (Myopia)|Myopia]] - **One** focal length
* Varifocal
  Incorporates all 3 powers (D,I,N - top-to-bottom), good for those with [[Presbyopia]]. Has [[Peripheral Distortion]]
* Bifocal
  Incorporates 2 powers (typically D and N), where one prescription takes up the majority of the lens and the second takes up a small segment of the lens
* Occupational
  Have 2 powers (I and N), good for computer and reading use
	`, datetime)

	if err != nil {
		fmt.Println("Failed to make new post", err)
	}

	posts, err := database.GetAllPosts()

	if err != nil {
		fmt.Println("Failed to get all posts", err)
		os.Exit(-1)
	}

	for _, post := range posts {
		fmt.Println(post)
	}

	// Print a message about the server starting
	fmt.Printf("Starting server on %s:%s\n", address, port)

	// log.Fatal() - If the server fails, log it
	// ListenAndServe() - Starts the server on the specified address and port
	// Sprintf allows strings to be formatted with variables
	log.Fatal(http.ListenAndServe(fmt.Sprintf("%s:%s", address, port), mux))
}
