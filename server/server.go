package server

import (
	"github.com/sirupsen/logrus"
	"go-spotify-cli/config"
	"go-spotify-cli/constants"
	"go-spotify-cli/handlers"
	"go-spotify-cli/routes"
	"go-spotify-cli/utils"
	"net/http"
)

func FetchAuthTokenFromBrowser() string {
	config.GlobalConfig.RequestedScopes = "user-modify-playback-state"
	BootstrapAuthServer(constants.AuthRoute)
	receivedToken := <-utils.AuthToken
	InitiateShutdown()
	return receivedToken
}

func FetchDeviceTokenFromBrowser() string {
	config.GlobalConfig.RequestedScopes = "user-read-playback-state"
	BootstrapAuthServer(constants.DeviceRoute)
	receivedToken := <-handlers.DeviceToken
	InitiateShutdown()
	return receivedToken
}

func GetAuthTokenOrFetchFromServer() string {
	token := utils.ReadJWTToken()
	if len(token) == 0 {
		token = FetchAuthTokenFromBrowser()
	}
	return token
}

var Shutdown = make(chan struct{})

func StartServer() {
	routes.SetupRoutes()

	server := &http.Server{Addr: constants.Port}

	// Start the server in a goroutine
	go func() {
		logrus.Info("Listening on " + constants.ServerUrl)
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
