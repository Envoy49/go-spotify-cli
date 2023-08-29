package routes

import (
	"go-spotify-cli/constants"
	"go-spotify-cli/handlers"
	"net/http"
)

func SetupRoutes() {
	http.HandleFunc(constants.AuthRoute, func(w http.ResponseWriter, r *http.Request) {
		handlers.StartAuthTokenFlow(w, r, constants.AuthCallBackRoute)
	})
	http.HandleFunc(constants.AuthCallBackRoute, handlers.FetchAccessToken)
	http.HandleFunc(constants.DeviceRoute, func(w http.ResponseWriter, r *http.Request) {
		handlers.StartDeviceTokenFlow(constants.DeviceCallBackRoute)
	})
	http.HandleFunc(constants.DeviceCallBackRoute, handlers.FetchDeviceToken)
}
