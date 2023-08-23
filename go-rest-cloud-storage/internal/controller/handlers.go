package controller

import (
	"encoding/json"
	"fmt"
	"net/http"
	"regexp"

	"github.com/mikouaj/go-rest-cloud-storage/internal/api"
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

func (c *controller) bucketHandler(w http.ResponseWriter, r *http.Request) {
	bucketName, err := getBucketNameFromPath(r.URL.Path)
	if err != nil {
		JSON(w, api.Error{Message: err.Error()}, http.StatusBadRequest)
		return
	}
	objects, err := c.storage.ListObjects(bucketName)
	if err != nil {
		JSON(w, api.Error{Message: err.Error()}, http.StatusInternalServerError)
		return
	}
	JSON(w, objects, http.StatusOK)
}

func getBucketNameFromPath(path string) (string, error) {
	regEx := regexp.MustCompile(fmt.Sprintf("^%s/([^/]+)$", bucketsPathPrefix))
	if !regEx.MatchString(path) {
		return "", fmt.Errorf("illegal bucket path %q", path)
	}
	matches := regEx.FindStringSubmatch(path)
	if len(matches) != 2 {
		return "", fmt.Errorf("illegal bucket path %q", path)
	}
	return matches[1], nil
}
