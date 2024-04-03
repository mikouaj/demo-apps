// Package controller provides implementation of REST controller
package controller

import (
	"fmt"
	"net"
	"net/http"

	"github.com/go-logr/logr"
	"github.com/prometheus/client_golang/prometheus"
	"github.com/prometheus/client_golang/prometheus/promauto"
	"github.com/prometheus/client_golang/prometheus/promhttp"
)

type controller struct {
	log            logr.Logger
	pathHandlers   map[string]http.Handler
	reqTotalMetric *prometheus.CounterVec
	connOpenMetric prometheus.Gauge
}

type auditedResponseWriter struct {
	http.ResponseWriter
	statusCode int
}

func (w *auditedResponseWriter) WriteHeader(code int) {
	w.statusCode = code
	w.ResponseWriter.WriteHeader(code)
}

func New() *controller {
	c := controller{
		pathHandlers: make(map[string]http.Handler),
		reqTotalMetric: promauto.NewCounterVec(prometheus.CounterOpts{
			Name: "go_rest_http_requests_total",
			Help: "The total number of HTTP requests",
		}, []string{"path", "status"}),
		connOpenMetric: promauto.NewGauge(prometheus.GaugeOpts{
			Name: "go_rest_http_connections",
			Help: "The number of active and idle connections"}),
	}
	c.RegisterPathHandler("/healthz", http.HandlerFunc(healthzHandler))
	c.RegisterPathHandler("/books", c.wrapHandlerWithAudit(http.HandlerFunc(booksHandler)))
	c.RegisterPathHandler("/metrics", promhttp.Handler())
	return &c
}

func (c *controller) SetLogger(log logr.Logger) {
	c.log = log
}

func (c *controller) RegisterPathHandler(path string, handler http.Handler) {
	c.pathHandlers[path] = handler
}

func (c *controller) ConnStateFunc(conn net.Conn, state http.ConnState) {
	switch state {
	case http.StateNew:
		c.connOpenMetric.Inc()
	case http.StateClosed, http.StateHijacked:
		c.connOpenMetric.Dec()
	}
}

func (c *controller) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	c.log.V(2).WithValues("remoteAddr", r.RemoteAddr).Info("Handling request")
	if h, ok := c.pathHandlers[r.URL.Path]; ok {
		h.ServeHTTP(w, r)
	} else {
		http.NotFound(w, r)
	}
}

func (c *controller) wrapHandlerWithAudit(handler http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		auditedWriter := &auditedResponseWriter{ResponseWriter: w}
		handler.ServeHTTP(auditedWriter, r)
		c.reqTotalMetric.With(prometheus.Labels{"path": r.URL.Path, "status": fmt.Sprint(auditedWriter.statusCode)}).Inc()
	})
}
