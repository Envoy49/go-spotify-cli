package routes

import (
	"go-spotify-cli/constants"
	"go-spotify-cli/handlers"
	"net/http"
	"sync"
)

var once sync.Once

func SetupRoutes() {
	once.Do(func() {
		http.HandleFunc(constants.UserModifyPlaybackStateRoute, func(w http.ResponseWriter, r *http.Request) {
			handlers.StartAuthTokenFlow(constants.UserModifyPlaybackStateRouteCallback)
		})
		http.HandleFunc(constants.UserModifyPlaybackStateRouteCallback, handlers.UserModifyTokenHandler)
		http.HandleFunc(constants.UserReadPlaybackStateRoute, func(w http.ResponseWriter, r *http.Request) {
			handlers.StartAuthTokenFlow(constants.UserReadPlaybackStateRouteCallback)
		})
		http.HandleFunc(constants.UserReadPlaybackStateRouteCallback, handlers.UserReadTokenHandler)
	})
}
