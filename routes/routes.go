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
		http.HandleFunc(constants.AuthRoute, func(w http.ResponseWriter, r *http.Request) {
			handlers.StartAuthTokenFlow(constants.AuthCallBackRoute)
		})
		http.HandleFunc(constants.AuthCallBackRoute, handlers.UserModifyTokenHandler)
		http.HandleFunc(constants.DeviceRoute, func(w http.ResponseWriter, r *http.Request) {
			handlers.StartAuthTokenFlow(constants.DeviceCallBackRoute)
		})
		http.HandleFunc(constants.DeviceCallBackRoute, handlers.UserReadTokenHandler)
	})
}
