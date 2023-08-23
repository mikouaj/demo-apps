package main

import (
	"context"
	"errors"
	"net/http"
	"os"
	"os/signal"
	"syscall"

	"github.com/go-logr/zapr"
	"github.com/mikouaj/go-rest-client/internal/controller"
	"github.com/mikouaj/go-rest-client/internal/provider"
	"go.uber.org/zap"
)

const (
	PortEnvVar     = "GO_REST_CLIENT_PORT"
	AppNameEnvVar  = "GO_REST_CLIENT_APP_NAME"
	DataURLEnvVar  = "GO_REST_CLIENT_DATA_URL"
	DefaultPort    = "8080"
	DefaultAppName = "go-rest-client"
)

func main() {
	zapConfig := zap.NewProductionConfig()
	zapConfig.Level = zap.NewAtomicLevelAt(-2)
	log := zapr.NewLogger(zap.Must(zapConfig.Build()))

	ctrl := controller.New(getDataProviderUsingEnvs())
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

func getDataProviderUsingEnvs() provider.DataProvider {
	var appName string
	if appName = os.Getenv(AppNameEnvVar); appName == "" {
		appName = DefaultAppName
	}
	if dataURL := os.Getenv(DataURLEnvVar); dataURL != "" {
		return provider.NewHTTPDataProvider(appName, dataURL)
	}
	return provider.NewLocalDataProvider(appName)
}
