package controller

import (
	"encoding/json"
	"net/http"

	"github.com/mikouaj/go-rest/internal/api"
)

func JSON(w http.ResponseWriter, in interface{}, code int) {
	w.Header().Set("Content-Type", "application/json; charset=utf-8")
	w.WriteHeader(code)
	json.NewEncoder(w).Encode(in)
}

func healthzHandler(w http.ResponseWriter, r *http.Request) {
	resp := api.Health{Status: "Up"}
	JSON(w, resp, http.StatusOK)
}

func booksHandler(w http.ResponseWriter, r *http.Request) {
	resp := api.Books{
		{
			Title:    "Ubik",
			Author:   "Philip K. Dick",
			Category: "Science Fiction",
		},
		{
			Title:    "Enders game",
			Author:   "Orson Scott Card",
			Category: "Science Fiction",
		},
	}
	JSON(w, resp, http.StatusOK)
}
