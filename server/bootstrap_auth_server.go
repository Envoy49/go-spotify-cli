package server

import (
	"go-spotify-cli/config"
	"go-spotify-cli/constants"
)

func FetchUserModifyTokenFromBrowser() string {
	config.GlobalConfig.RequestedScopes = "user-modify-playback-state"
	BootstrapAuthServer(constants.AuthRoute)
	receivedToken := <-config.AuthTokenData
	InitiateShutdown()
	return receivedToken.UserModifyToken
}

func FetchUserReadTokenFromBrowser() string {
	config.GlobalConfig.RequestedScopes = "user-read-playback-state"
	BootstrapAuthServer(constants.DeviceRoute)
	receivedToken := <-config.AuthTokenData
	InitiateShutdown()
	return receivedToken.UserReadToken
}

func ReadUserModifyTokenOrFetchFromServer() string {
	token := config.ReadTokenFromHome("userModifyToken")
	if len(token) == 0 {
		token = FetchUserModifyTokenFromBrowser()
	}
	return token
}

func ReadUserReadTokenOrFetchFromServer() string {
	token := config.ReadTokenFromHome("userReadToken")
	if len(token) == 0 {
		token = FetchUserReadTokenFromBrowser()
	}
	return token
}
