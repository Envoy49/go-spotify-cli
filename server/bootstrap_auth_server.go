package server

import (
	"go-spotify-cli/config"
	"go-spotify-cli/constants"
)

func FetchUserModifyTokenFromBrowser() string {
	config.GlobalConfig.RequestedScopes = constants.UserModifyPlaybackStateScope
	BootstrapAuthServer(constants.UserModifyPlaybackStateRoute)
	receivedToken := <-config.AuthTokenData
	InitiateShutdown()
	return receivedToken.UserModifyToken
}

func FetchUserReadTokenFromBrowser() string {
	config.GlobalConfig.RequestedScopes = constants.UserReadPlaybackState
	BootstrapAuthServer(constants.UserReadPlaybackStateRoute)
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
