package routes

import (
	"github.com/envoy49/go-spotify-cli/config"
	"github.com/envoy49/go-spotify-cli/handlers"
	"net/http"
	"sync"
)

var once sync.Once

func TokenHandlerWithParams(cfg *config.Config, tokenType config.TokenType) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		handlers.TokenHandler(w, r, cfg, tokenType)
	}
}

func SetupRoutes(cfg *config.Config) {
	once.Do(func() {
		http.HandleFunc(config.UserModifyPlaybackStateRoute, func(w http.ResponseWriter, r *http.Request) {
			handlers.StartAuthTokenFlow(cfg, config.UserModifyPlaybackStateRouteCallback)
		})
		http.HandleFunc(config.UserReadPlaybackStateRoute, func(w http.ResponseWriter, r *http.Request) {
			handlers.StartAuthTokenFlow(cfg, config.UserReadPlaybackStateRouteCallback)
		})
		http.HandleFunc(config.UserLibraryReadRoute, func(w http.ResponseWriter, r *http.Request) {
			handlers.StartAuthTokenFlow(cfg, config.UserLibraryReadRouteCallback)
		})
		// Callback routes
		http.HandleFunc(config.UserModifyPlaybackStateRouteCallback, TokenHandlerWithParams(cfg, config.ModifyToken))
		http.HandleFunc(config.UserReadPlaybackStateRouteCallback, TokenHandlerWithParams(cfg, config.ReadToken))
		http.HandleFunc(config.UserLibraryReadRouteCallback, TokenHandlerWithParams(cfg, config.LibraryRead))
	})
}
