package main

import (
	"fmt"
	_ "github.com/joho/godotenv/autoload" // Used to autoload .env files
	"html"
	"log"
	"net/http"
	"os"
)

func main() {
	address := os.Getenv("ADDRESS")
	port := os.Getenv("PORT")

	http.HandleFunc("/api", handleRoot)

	log.Fatal(http.ListenAndServe(fmt.Sprintf("%s:%s", address, port), nil))
}

func handleRoot(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Hello, %q", html.EscapeString(r.URL.Path))
}
