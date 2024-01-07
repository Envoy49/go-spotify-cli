package routes

import (
	"github.com/envoy49/go-spotify-cli/config"
	"github.com/envoy49/go-spotify-cli/handlers"
	"net/http"
	"sync"
)

var once sync.Once

func TokenHandlerWithParams(tokenType config.TokenType) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		handlers.TokenHandler(w, r, tokenType)
	}
}

func SetupRoutes() {
	once.Do(func() {
		http.HandleFunc(config.UserModifyPlaybackStateRoute, func(w http.ResponseWriter, r *http.Request) {
			handlers.StartAuthTokenFlow(config.UserModifyPlaybackStateRouteCallback)
		})
		http.HandleFunc(config.UserReadPlaybackStateRoute, func(w http.ResponseWriter, r *http.Request) {
			handlers.StartAuthTokenFlow(config.UserReadPlaybackStateRouteCallback)
		})
		http.HandleFunc(config.UserLibraryReadRoute, func(w http.ResponseWriter, r *http.Request) {
			handlers.StartAuthTokenFlow(config.UserLibraryReadRouteCallback)
		})
		// Callback routes
		http.HandleFunc(config.UserModifyPlaybackStateRouteCallback, TokenHandlerWithParams(config.ModifyToken))
		http.HandleFunc(config.UserReadPlaybackStateRouteCallback, TokenHandlerWithParams(config.ReadToken))
		http.HandleFunc(config.UserLibraryReadRouteCallback, TokenHandlerWithParams(config.LibraryRead))
	})
}
