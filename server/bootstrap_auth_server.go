package server

import (
	"github.com/sirupsen/logrus"
	"go-spotify-cli/config"
	"net/http"
)

func BootstrapAuthServer(route string) {
	// Start the server in a goroutine
	go StartServer()

	resp, err := http.Get("http://localhost" + config.GlobalConfig.Port + route)
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
