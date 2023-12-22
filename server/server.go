package server

import (
	"context"
	"github.com/sirupsen/logrus"
	"go-spotify-cli/constants"
	"go-spotify-cli/routes"
	"net/http"
	"time"
)

var Shutdown = make(chan struct{})

func StartServer() {
	routes.SetupRoutes()

	server := &http.Server{Addr: constants.Port}

	// Start the server in a goroutine
	go func() {
		logrus.Println("Opened server to get an auth token on " + constants.ServerUrl)
		if err := server.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			logrus.WithError(err).Error("Error starting the server")
		}
	}()

	// Listen for a signal to shut down
	<-Shutdown

	// Create a deadline to wait for.
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	// Shutdown the server gracefully
	if err := server.Shutdown(ctx); err != nil {
		logrus.WithError(err).Error("Error shutting down the server")
	}
}

func InitiateShutdown() {
	Shutdown <- struct{}{}
}

func BootstrapAuthServer(route string) {
	// Start the server in a goroutine
	go StartServer()

	resp, err := http.Get(constants.ServerUrl + route)
	if err != nil {
		logrus.WithError(err).Error("Error making the GET request to: " + route)
		return
	}

	defer func() {
		err := resp.Body.Close()
		if err != nil {
			logrus.WithError(err).Error("Error closing request to :" + route)
		}
	}()
}
