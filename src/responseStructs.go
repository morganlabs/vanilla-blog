package main

// An object, essentially
// Defines a Struct, with keys and types
type Alive struct {
	Alive  bool   `json:"alive"`
	Uptime string `json:"uptime"`
}

type ApiError struct {
	Status  int    `json:"status"`
	Message string `json:"message"`
}
