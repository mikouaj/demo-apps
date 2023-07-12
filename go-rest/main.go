package main

import (
	"context"
	"errors"
	"net/http"
	"os"
	"os/signal"
	"syscall"

	"github.com/go-logr/zapr"
	"github.com/mikouaj/go-rest/internal/controller"
	"go.uber.org/zap"
)

func main() {
	zapConfig := zap.NewProductionConfig()
	zapConfig.Level = zap.NewAtomicLevelAt(-2)
	log := zapr.NewLogger(zap.Must(zapConfig.Build()))

	ctrl := controller.New()
	ctrl.SetLogger(log)
	srv := &http.Server{
		Addr:    ":8080",
		Handler: ctrl,
	}

	go func() {
		log.Info("Starting HTTP server")
		if err := srv.ListenAndServe(); err != nil && !errors.Is(err, http.ErrServerClosed) {
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
