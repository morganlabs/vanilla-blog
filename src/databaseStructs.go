package main

type PostSummary struct {
	ID          string `json:"id"`
	Title       string `json:"title"`
	Summary     string `json:"summary"`
	PublishedAt string `json:"publishedAt"`
}

type Post struct {
	PostSummary
	Content string `json:"content"`
}
