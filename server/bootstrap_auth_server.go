package server

import (
	"github.com/sirupsen/logrus"
	"go-spotify-cli/constants"
	"net/http"
)

func BootstrapAuthServer(route string) {
	// Start the server in a goroutine
	go StartServer()

	resp, err := http.Get(constants.ServerUrl + route)
	if err != nil {
		logrus.WithError(err).Error("Error making the GET request for /auth route")
		return
	}

	defer func() {
		err := resp.Body.Close()
		if err != nil {
			logrus.WithError(err).Error("Error closing request for /auth")
		}
	}()
}
