package main

import (
	"fmt"
	"go-spotify-cli/cmd/player"
	"go-spotify-cli/config"
	"go-spotify-cli/handlers"
	"go-spotify-cli/utils"
	"net/http"
	"time"
)

func init() {
	config.LoadConfiguration()
	
}

func main() {
	http.HandleFunc("/auth", handlers.StartAuthentication)  // This will initiate the authentication
	http.HandleFunc("/callback", handlers.FetchAccessToken) // This will handle the callback after authentication

	// Start the server in a goroutine to allow further execution
	go func() {
		fmt.Printf("Listening on %s\n", config.GlobalConfig.ServerUrl)
		if err := http.ListenAndServe(config.GlobalConfig.Port, nil); err != nil {
			fmt.Println("Error listening starting the server", err)
		}
	}()

	// Allow the server some time to start
	time.Sleep(2 * time.Second) // You can adjust the sleep duration

	var token = utils.ReadJWTToken()

	if len(token) == 0 {
		// Make a request to the server to initiate authentication
		resp, err := http.Get("http://localhost" + config.GlobalConfig.Port + "/auth")
		if err != nil {
			fmt.Println("Error making the GET request:", err)
			return
		}
		defer resp.Body.Close()
		// Handle the response if needed.
		fmt.Println("Response status:", resp.Status)
	} else {
		if playErr := player.Play(token); playErr != nil {
			utils.PrintError("Failed to get Play your track:", playErr)
		}
	}

	// The main function will keep running because of the server goroutine
	select {} // Keep the main function running indefinitely
}
