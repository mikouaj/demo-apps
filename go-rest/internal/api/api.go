// Package api provides structs that define application's REST API
package api

type Health struct {
	Status string `json:"status"`
}

type Book struct {
	Title    string `json:"title,omitempty"`
	Author   string `json:"author,omitempty"`
	Category string `json:"category,omitempty"`
}

type Books []Book
