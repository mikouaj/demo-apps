// Package controller provides implementation of REST controller
package controller

import (
	"net/http"

	"github.com/go-logr/logr"
)

type controller struct {
	log          logr.Logger
	pathHandlers map[string]http.HandlerFunc
}

func New() *controller {
	return &controller{
		pathHandlers: map[string]http.HandlerFunc{
			"/healthz": healthzHandler,
			"/books":   booksHandler,
		},
	}
}

func (c *controller) SetLogger(log logr.Logger) {
	c.log = log
}

func (c *controller) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	c.log.V(2).WithValues("remoteAddr", r.RemoteAddr).Info("Handling request")
	handlerFunc := http.NotFound
	if f, ok := c.pathHandlers[r.URL.Path]; ok {
		handlerFunc = f
	}
	handlerFunc(w, r)
}
