package server

import (
	"context"
	"net/http"
	"time"

	"github.com/envoy49/go-spotify-cli/config"

	"github.com/envoy49/go-spotify-cli/routes"
	"github.com/sirupsen/logrus"
)

const (
	serverPort = ":4949"
)

func Server(ctx context.Context, cfg *config.Config) {
	// Create a new server instance each time
	server := &http.Server{Addr: serverPort}

	routes.SetupRoutes(cfg)

	// Start the server in a goroutine
	go func() {
		logrus.Println("Opened server to get an auth token on " + config.ServerUrl)
		if err := server.ListenAndServe(); err != http.ErrServerClosed {
			logrus.WithError(err).Error("Error starting the server")
		}
	}()

	// Listen for the context being cancelled
	<-ctx.Done()

	// Create a deadline to wait for
	shutdownCtx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	// Shutdown the server gracefully
	if err := server.Shutdown(shutdownCtx); err != nil {
		logrus.WithError(err).Error("Error shutting down the server")
	}
}

func StartServer(cfg *config.Config, route string) context.CancelFunc {
	ctx, cancel := context.WithCancel(context.Background())
	go Server(ctx, cfg)

	resp, err := http.Get(config.ServerUrl + route)
	if err != nil {
		logrus.WithError(err).Error("Error making the GET request to: " + route)
		return cancel
	}

	defer func() {
		err := resp.Body.Close()
		if err != nil {
			logrus.WithError(err).Error("Error closing request to :" + route)
		}
	}()

	return cancel
}
