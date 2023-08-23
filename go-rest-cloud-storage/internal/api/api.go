// Package api provides structs that define application's REST API
package api

type Health struct {
	Status string `json:"status"`
}

type StorageObject struct {
	Name string `json:"name,omitempty"`
}

type StorageObjects []StorageObject

type Error struct {
	Message string `json:"message,omitempty"`
}
