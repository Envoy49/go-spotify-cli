package server

import (
	"fmt"
	"go-spotify-cli/config"
	"net/http"
)

func BootstrapAuthServer(route string) {
	// Start the server in a goroutine
	go StartServer()

	resp, err := http.Get("http://localhost" + config.GlobalConfig.Port + route)
	if err != nil {
		fmt.Println("Error making the GET request for /auth route:", err)
		return
	}
	defer func() {
		err := resp.Body.Close()
		if err != nil {
			fmt.Println("Error closing request for /auth", err)
		}
	}()
}
