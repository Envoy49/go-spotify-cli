package server

import (
	"github.com/sirupsen/logrus"
	"go-spotify-cli/constants"
	"go-spotify-cli/routes"
	"net/http"
)

var Shutdown = make(chan struct{})

func StartServer() {
	routes.SetupRoutes()

	server := &http.Server{Addr: constants.Port}

	// Start the server in a goroutine
	go func() {
		logrus.Info("Opened server to get an auth token on " + constants.ServerUrl)
		if err := server.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			// This will print if the server is forcibly closed.
			logrus.WithError(err).Error("Error starting the server")

		}
	}()

	// Run server until we receive a shutdown signal
	<-Shutdown

	// Now, gracefully shut down the server
	if err := server.Close(); err != nil {
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
