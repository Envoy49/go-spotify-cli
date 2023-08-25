package routes

import (
	"go-spotify-cli/handlers"
	"net/http"
)

func SetupRoutes() {
	http.HandleFunc("/auth", handlers.StartAuthentication)
	http.HandleFunc("/callback", handlers.FetchAccessToken)
}
