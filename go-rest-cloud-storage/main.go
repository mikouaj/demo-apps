package main

import (
	"context"
	"errors"
	"net/http"
	"os"
	"os/signal"
	"syscall"

	"github.com/go-logr/zapr"
	"github.com/mikouaj/go-rest-cloud-storage/internal/controller"
	"github.com/mikouaj/go-rest-cloud-storage/storage"
	"go.uber.org/zap"
)

const (
	PortEnvVar  = "GO_REST_CLIENT_PORT"
	DefaultPort = "8080"
)

func main() {
	zapConfig := zap.NewProductionConfig()
	zapConfig.Level = zap.NewAtomicLevelAt(-2)
	log := zapr.NewLogger(zap.Must(zapConfig.Build()))

	ctx := context.Background()
	storage, err := storage.NewStorageClient(ctx)
	if err != nil {
		log.Error(err, "unable to create Cloud Storage client")
		return
	}
	ctrl := controller.New(storage)
	ctrl.SetLogger(log)
	addr := ":" + DefaultPort
	if port := os.Getenv(PortEnvVar); port != "" {
		addr = ":" + port
	}
	srv := &http.Server{
		Addr:    addr,
		Handler: ctrl,
	}

	go func() {
		log.Info("Starting HTTP server")
		if err := srv.ListenAndServe(); err != nil && !errors.Is(err, http.ErrServerClosed) {
			log.Error(err, "http server error")
		}
	}()
	ctx, cancel := context.WithCancel(ctx)
	defer cancel()
	shutdown := make(chan os.Signal, 1)
	signal.Notify(shutdown, syscall.SIGINT, syscall.SIGTERM)
	<-shutdown
	log.Info("Stopping HTTP server")
	if err := srv.Shutdown(ctx); err != nil {
		log.Error(err, "http server error")
	}
}
