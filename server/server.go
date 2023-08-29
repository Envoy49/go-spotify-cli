package server

import (
	"fmt"
	"go-spotify-cli/config"
	"go-spotify-cli/constants"
	"go-spotify-cli/handlers"
	"go-spotify-cli/routes"
	"go-spotify-cli/utils"
	"net/http"
)

func FetchAuthTokenFromBrowser() string {
	BootstrapAuthServer(constants.AuthRoute)
	receivedToken := <-utils.AuthToken
	InitiateShutdown()
	return receivedToken
}

func FetchDeviceTokenFromBrowser() string {
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

	server := &http.Server{Addr: config.GlobalConfig.Port}

	// Start the server in a goroutine
	go func() {
		fmt.Printf("Listening on %s\n", config.GlobalConfig.ServerUrl)
		if err := server.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			// This will print if the server is forcibly closed.
			fmt.Println("Error starting the server:", err)
		}
	}()

	// Run server until we receive a shutdown signal
	<-Shutdown

	// Now, gracefully shut down the server
	if err := server.Close(); err != nil {
		fmt.Println("Error shutting down the server:", err)
	}
}

func InitiateShutdown() {
	Shutdown <- struct{}{}
}
