package routes

import (
	"go-spotify-cli/constants"
	"go-spotify-cli/handlers"
	"net/http"
	"sync"
)

var once sync.Once

func TokenHandlerWithParams(tokenType constants.TokenType) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		handlers.TokenHandler(w, r, tokenType)
	}
}

func SetupRoutes() {
	once.Do(func() {
		http.HandleFunc(constants.UserModifyPlaybackStateRoute, func(w http.ResponseWriter, r *http.Request) {
			handlers.StartAuthTokenFlow(constants.UserModifyPlaybackStateRouteCallback)
		})
		http.HandleFunc(constants.UserReadPlaybackStateRoute, func(w http.ResponseWriter, r *http.Request) {
			handlers.StartAuthTokenFlow(constants.UserReadPlaybackStateRouteCallback)
		})
		http.HandleFunc(constants.UserLibraryReadRoute, func(w http.ResponseWriter, r *http.Request) {
			handlers.StartAuthTokenFlow(constants.UserLibraryReadRouteCallback)
		})
		// Callback routes
		http.HandleFunc(constants.UserModifyPlaybackStateRouteCallback, TokenHandlerWithParams(constants.ModifyToken))
		http.HandleFunc(constants.UserReadPlaybackStateRouteCallback, TokenHandlerWithParams(constants.ReadToken))
		http.HandleFunc(constants.UserLibraryReadRouteCallback, TokenHandlerWithParams(constants.LibraryRead))
	})
}
