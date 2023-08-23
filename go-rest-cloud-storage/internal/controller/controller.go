// Package controller provides implementation of REST controller
package controller

import (
	"net/http"
	"strings"

	"github.com/go-logr/logr"
	"github.com/mikouaj/go-rest-cloud-storage/storage"
)

const (
	bucketsPathPrefix = "/buckets"
)

type controller struct {
	log                logr.Logger
	pathHandlers       map[string]http.HandlerFunc
	pathPrefixHandlers map[string]http.HandlerFunc
	storage            storage.StorageClient
}

func New(storage storage.StorageClient) *controller {
	c := &controller{
		storage: storage,
	}
	c.pathHandlers = map[string]http.HandlerFunc{
		"/healthz": c.healthzHandler,
	}
	c.pathPrefixHandlers = map[string]http.HandlerFunc{
		bucketsPathPrefix: c.bucketHandler,
	}
	return c
}

func (c *controller) SetLogger(log logr.Logger) {
	c.log = log
}

func (c *controller) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	handlerFunc := http.NotFound
	if f, ok := c.pathHandlers[r.URL.Path]; ok {
		handlerFunc = f
	}
	for p, f := range c.pathPrefixHandlers {
		if strings.HasPrefix(r.URL.Path, p) {
			handlerFunc = f
		}
	}
	handlerFunc(w, r)
}
