package server

import (
	"fmt"
	"go-spotify-cli/config"
	"net/http"
)

func StartAuthentication() {
	// Start the server in a goroutine
	go StartServer()

	resp, err := http.Get("http://localhost" + config.GlobalConfig.Port + "/auth")
	if err != nil {
		fmt.Println("Error making the GET request:", err)
		return
	}
	defer func() {
		err := resp.Body.Close()
		if err != nil {
			fmt.Println("Error closing request for /auth", err)
		}
	}()

}
