package controller

import (
	"encoding/json"
	"net/http"

	"github.com/mikouaj/go-rest-client/internal/api"
)

func JSON(w http.ResponseWriter, in interface{}, code int) {
	w.Header().Set("Content-Type", "application/json; charset=utf-8")
	w.WriteHeader(code)
	json.NewEncoder(w).Encode(in)
}

func (c *controller) healthzHandler(w http.ResponseWriter, r *http.Request) {
	resp := api.Health{Status: "Up"}
	JSON(w, resp, http.StatusOK)
}

func (c *controller) dataHandler(w http.ResponseWriter, r *http.Request) {
	data, err := c.dataProvider.GetData(r.Context())
	if err != nil {
		JSON(w, api.Error{Message: err.Error()}, http.StatusInternalServerError)
	}
	JSON(w, data, http.StatusOK)
}
