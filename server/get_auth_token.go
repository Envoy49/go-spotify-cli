package server

import (
	"go-spotify-cli/auth"
	"go-spotify-cli/config"
	"go-spotify-cli/constants"
)

func ReadUserModifyTokenOrFetchFromServer() string {
	tokenInstance := config.ReadTokenFromHome("userModifyToken")
	if len(tokenInstance.UserModifyRefreshToken) > 0 {
		newToken, err := auth.RefreshAuthToken(tokenInstance.UserModifyRefreshToken)
		if err != nil {
			return FetchUserModifyTokenFromBrowser()
		}

		userModifyToken := config.TokenStructure{
			UserModifyToken:          newToken.AccessToken,
			UserModifyTokenExpiresIn: int64(newToken.ExpiresIn),
		}
		config.WriteTokenToHomeDirectory(&userModifyToken, false)

		return newToken.AccessToken
	}

	if len(tokenInstance.UserModifyToken) == 0 {
		return FetchUserModifyTokenFromBrowser()
	}

	return tokenInstance.UserModifyToken
}

func ReadUserReadTokenOrFetchFromServer() string {
	tokenInstance := config.ReadTokenFromHome("userReadToken")
	if len(tokenInstance.UserReadRefreshToken) > 0 {
		newToken, err := auth.RefreshAuthToken(tokenInstance.UserReadRefreshToken)

		if err != nil {
			return FetchUserReadTokenFromBrowser()
		}

		userReadToken := config.TokenStructure{
			UserReadToken:          newToken.AccessToken,
			UserReadTokenExpiresIn: int64(newToken.ExpiresIn),
		}
		config.WriteTokenToHomeDirectory(&userReadToken, false)
		return newToken.AccessToken
	}

	if len(tokenInstance.UserReadToken) == 0 {
		return FetchUserReadTokenFromBrowser()
	}
	return tokenInstance.UserReadToken
}

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
