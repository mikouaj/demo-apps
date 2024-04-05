package main

import (
	"context"
	"errors"
	"net/http"
	"os"
	"os/signal"
	"syscall"

	"github.com/go-logr/zapr"
	"github.com/mikouaj/go-rest/internal/config"
	"github.com/mikouaj/go-rest/internal/controller"
	"go.uber.org/zap"
)

func main() {
	zapConfig := zap.NewProductionConfig()
	zapConfig.Level = zap.NewAtomicLevelAt(-2)
	log := zapr.NewLogger(zap.Must(zapConfig.Build()))

	config, err := config.LoadFromEnv()
	if err != nil {
		log.Error(err, "failed to load configuration")
	}
	ctrl := controller.New()
	ctrl.SetLogger(log)
	srv := &http.Server{
		Addr:      ":8080",
		Handler:   ctrl,
		ConnState: ctrl.ConnStateFunc,
	}

	go func() {
		var err error
		if config.ListenTLS {
			log.Info("Starting HTTPS server")
			srv.Addr = ":8443"
			err = srv.ListenAndServeTLS(config.TLSCertPath, config.TLSKeyPath)
		} else {
			log.Info("Starting HTTP server")
			err = srv.ListenAndServe()
		}
		if err != nil && !errors.Is(err, http.ErrServerClosed) {
			log.Error(err, "http server error")
		}
	}()
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()
	shutdown := make(chan os.Signal, 1)
	signal.Notify(shutdown, syscall.SIGINT, syscall.SIGTERM)
	<-shutdown
	log.Info("Stopping HTTP server")
	if err := srv.Shutdown(ctx); err != nil {
		log.Error(err, "http server error")
	}
}
