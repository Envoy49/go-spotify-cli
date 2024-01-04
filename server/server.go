package server

import (
	"context"
	"github.com/sirupsen/logrus"
	"github.com/envoy49/go-spotify-cli/constants"
	"github.com/envoy49/go-spotify-cli/routes"
	"net/http"
	"time"
)

func Server(ctx context.Context) {
	// Create a new server instance each time
	server := &http.Server{Addr: constants.Port}

	routes.SetupRoutes()

	// Start the server in a goroutine
	go func() {
		logrus.Println("Opened server to get an auth token on " + constants.ServerUrl)
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

func StartServer(route string) context.CancelFunc {
	ctx, cancel := context.WithCancel(context.Background())
	go Server(ctx)

	resp, err := http.Get(constants.ServerUrl + route)
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
