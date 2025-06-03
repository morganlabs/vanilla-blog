package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"sort"
	"time"
)

func getAllPosts(w http.ResponseWriter, r *http.Request) {
	posts, err := database.GetAllPosts()

	if err != nil {
		writeError(w, 500, "Failed to get posts")
		fmt.Println(err)
		return
	}

	sort.Slice(posts, func(i, j int) bool {
		t1, _ := time.Parse("Mon Jan _2 15:04:05 2006", posts[i].PublishedAt)
		t2, _ := time.Parse("Mon Jan _2 15:04:05 2006", posts[j].PublishedAt)
		return t1.After(t2)
	})

	writeJSON(w, 200, posts)
}

func getPostFromID(w http.ResponseWriter, r *http.Request) {
	id := r.PathValue("id")                          // Get the query param "id"
	toHTML := r.URL.Query().Get("toHTML") != "false" // Always convert MD to HTML by default
	post, err := database.GetPost(id, toHTML)

	if err != nil {
		writeError(w, 500, fmt.Sprintf("Failed to get post with ID %s", id))
		fmt.Println(err)
		return
	}

	writeJSON(w, 200, post)
}

func makeNewPost(w http.ResponseWriter, r *http.Request) {
	var post Post

	err := json.NewDecoder(r.Body).Decode(&post)
	if err != nil {
		fmt.Println(err)
		writeError(w, 500, "Failed to encode JSON from Request Body")
	}

	defer r.Body.Close()

	id, err := database.NewPost(post.Title, post.Summary, post.Content, post.PublishedAt)
	if err != nil {
		fmt.Println(err)
		writeError(w, 500, "Failed to create new Post (Write)")
	}

	newPost, err := database.GetPost(id, false)
	if err != nil {
		fmt.Println(err)
		writeError(w, 500, "Failed to create new Post (Read)")
	}

	writeJSON(w, 200, newPost)
}
