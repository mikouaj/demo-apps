// Package controller provides implementation of REST controller
package controller

import (
	"net/http"

	"github.com/go-logr/logr"
	"github.com/mikouaj/go-rest-client/internal/provider"
)

type controller struct {
	log          logr.Logger
	pathHandlers map[string]http.HandlerFunc
	dataProvider provider.DataProvider
}

func New(dp provider.DataProvider) *controller {
	c := &controller{
		dataProvider: dp,
	}
	c.pathHandlers = map[string]http.HandlerFunc{
		"/healthz": c.healthzHandler,
		"/data":    c.dataHandler,
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
	handlerFunc(w, r)
}
