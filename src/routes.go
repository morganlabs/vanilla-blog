package main

import (
	"fmt"
	"net/http"
	"time"
)

type MethodHandler = func(w http.ResponseWriter, r *http.Request)

func registerRoutes(mux *http.ServeMux) {
	registerRoute(mux, "/api", apiRoot, notAllowed, notAllowed)

	// Post-related
	registerRoute(mux, "/api/posts", getAllPosts, makeNewPost, notAllowed)
	registerRoute(mux, "/api/posts/{id}", getPostFromID, notAllowed, notAllowed)

	// Begin serving the website
	// This is so amazingly easy i love it
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

func registerRoute(mux *http.ServeMux, endpoint string, getMethod, postMethod, deleteMethod MethodHandler) {
	mux.HandleFunc(endpoint, func(w http.ResponseWriter, r *http.Request) {
		switch method := r.Method; method {
		case "GET":
			getMethod(w, r)
			return
		case "POST":
			postMethod(w, r)
			return
		case "DELETE":
			deleteMethod(w, r)
			return
		}
	})
}

func notAllowed(w http.ResponseWriter, r *http.Request) {
	writeError(w, 405, fmt.Sprintf("Method %s Not Allowed", r.Method))
}
