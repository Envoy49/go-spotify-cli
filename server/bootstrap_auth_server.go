package server

import (
	"go-spotify-cli/auth"
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
	authToken, refreshToken := config.ReadTokenFromHome("userModifyToken")
	if len(refreshToken) > 0 {
		newToken, err := auth.RefreshAuthToken(refreshToken)
		if err != nil {
			return FetchUserModifyTokenFromBrowser()
		}

		userModifyToken := config.TokenStructure{
			UserModifyToken:          newToken.AccessToken,
			UserModifyTokenExpiresIn: int64(newToken.ExpiresIn),
		}
		config.WriteTokenToHomeDirectory(&userModifyToken, false)

		authToken = newToken.AccessToken
	}

	if len(authToken) == 0 {
		return FetchUserModifyTokenFromBrowser()
	}

	return authToken
}

func ReadUserReadTokenOrFetchFromServer() string {
	authToken, refreshToken := config.ReadTokenFromHome("userReadToken")
	if len(refreshToken) > 0 {
		newToken, err := auth.RefreshAuthToken(refreshToken)

		if err != nil {
			return FetchUserReadTokenFromBrowser()
		}

		userReadToken := config.TokenStructure{
			UserReadToken:          newToken.AccessToken,
			UserReadTokenExpiresIn: int64(newToken.ExpiresIn),
		}
		config.WriteTokenToHomeDirectory(&userReadToken, false)
		authToken = newToken.AccessToken
	}

	if len(authToken) == 0 {
		return FetchUserReadTokenFromBrowser()
	}
	return authToken
}
