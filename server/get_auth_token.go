package server

import (
	"go-spotify-cli/auth"
	"go-spotify-cli/config"
	"go-spotify-cli/constants"
)

func ReadUserModifyTokenOrFetchFromServer() string {
	tokenInstance := config.ReadTokenFromHome(constants.ModifyToken)
	if len(tokenInstance.ModifyToken.UserModifyRefreshToken) > 0 {
		newToken, err := auth.RefreshAuthToken(tokenInstance.ModifyToken.UserModifyRefreshToken)
		if err != nil {
			return FetchUserModifyTokenFromBrowser()
		}

		userModifyToken := config.CombinedTokenStructure{
			ModifyToken: config.UserModifyTokenStructure{
				UserModifyToken:          newToken.AccessToken,
				UserModifyTokenExpiresIn: int64(newToken.ExpiresIn),
			},
		}
		config.WriteTokenToHomeDirectory(&userModifyToken, false)

		return newToken.AccessToken
	}

	if len(tokenInstance.ModifyToken.UserModifyToken) == 0 {
		return FetchUserModifyTokenFromBrowser()
	}

	return tokenInstance.ModifyToken.UserModifyToken
}

func ReadUserReadTokenOrFetchFromServer() string {
	tokenInstance := config.ReadTokenFromHome(constants.ReadToken)
	if len(tokenInstance.ReadToken.UserReadRefreshToken) > 0 {
		newToken, err := auth.RefreshAuthToken(tokenInstance.ReadToken.UserReadRefreshToken)

		if err != nil {
			return FetchUserReadTokenFromBrowser()
		}

		userReadToken := config.CombinedTokenStructure{
			ReadToken: config.UserReadTokenStructure{
				UserReadToken:          newToken.AccessToken,
				UserReadTokenExpiresIn: int64(newToken.ExpiresIn),
			},
		}
		config.WriteTokenToHomeDirectory(&userReadToken, false)
		return newToken.AccessToken
	}

	if len(tokenInstance.ReadToken.UserReadToken) == 0 {
		return FetchUserReadTokenFromBrowser()
	}
	return tokenInstance.ReadToken.UserReadToken
}

func FetchUserModifyTokenFromBrowser() string {
	config.GlobalConfig.RequestedScopes = constants.UserModifyPlaybackStateScope
	BootstrapAuthServer(constants.UserModifyPlaybackStateRoute)
	receivedToken := <-config.AuthTokenData
	InitiateShutdown()
	return receivedToken.ModifyToken.UserModifyToken
}

func FetchUserReadTokenFromBrowser() string {
	config.GlobalConfig.RequestedScopes = constants.UserReadPlaybackState
	BootstrapAuthServer(constants.UserReadPlaybackStateRoute)
	receivedToken := <-config.AuthTokenData
	InitiateShutdown()
	return receivedToken.ReadToken.UserReadToken
}
