package main

import (
	"fmt"
	"net/http"
	"time"
)

func registerRoutes(mux *http.ServeMux) {
	mux.HandleFunc("/api", apiRoot)
	mux.HandleFunc("/api/posts", getAllPosts)
	mux.HandleFunc("/api/posts/{id}", getPostFromID)
	mux.Handle("/", http.FileServer(http.Dir("./src/public")))
}

func apiRoot(w http.ResponseWriter, r *http.Request) {
	// Record the time difference between `startTime` and now
	// formatDuration() can be found below
	uptime := formatDuration(time.Since(startTime))

	// Define an instance of the Alive struct
	res := Alive{Alive: true, Uptime: uptime}

	writeJSON(w, 200, res)
}

func getAllPosts(w http.ResponseWriter, r *http.Request) {
	posts, err := database.GetAllPosts()

	if err != nil {
		writeError(w, 500, "Failed to get posts")
		fmt.Println(err)
		return
	}

	writeJSON(w, 200, posts)
}

func getPostFromID(w http.ResponseWriter, r *http.Request) {
	id := r.URL.Path[len("/api/posts/"):]            // Extract the URL Slug from the path
	toHTML := r.URL.Query().Get("toHTML") != "false" // Always convert MD to HTML by default
	post, err := database.GetPost(id, toHTML)

	if err != nil {
		writeError(w, 500, fmt.Sprintf("Failed to get post with ID %s", id))
		fmt.Println(err)
		return
	}

	writeJSON(w, 200, post)
}
