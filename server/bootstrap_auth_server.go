package server

import (
	"fmt"
	"go-spotify-cli/auth"
	"go-spotify-cli/config"
	"go-spotify-cli/constants"
)

func FetchUserModifyTokenFromBrowser(token string) string {
	if len(token) > 0 {
		return token
	}
	config.GlobalConfig.RequestedScopes = constants.UserModifyPlaybackStateScope
	BootstrapAuthServer(constants.UserModifyPlaybackStateRoute)
	receivedToken := <-config.AuthTokenData
	InitiateShutdown()
	return receivedToken.UserModifyToken
}

func FetchUserReadTokenFromBrowser(token string) string {
	if len(token) > 0 {
		return token
	}
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
			authToken = FetchUserModifyTokenFromBrowser("")
		}

		userModifyToken := config.TokenStructure{
			UserModifyToken:          newToken.AccessToken,
			UserModifyTokenExpiresIn: int64(newToken.ExpiresIn),
		}
		config.WriteTokenToHomeDirectory(&userModifyToken, false)

		authToken = newToken.AccessToken
	}

	if len(authToken) == 0 {
		authToken = FetchUserModifyTokenFromBrowser("")
	}

	return authToken
}

func ReadUserReadTokenOrFetchFromServer() string {
	authToken, refreshToken := config.ReadTokenFromHome("userReadToken")
	fmt.Println("---------->ReadToken auth", authToken)
	fmt.Println("---------->ReadToken refresh", refreshToken)

	if len(refreshToken) > 0 {
		newToken, err := auth.RefreshAuthToken(refreshToken)

		if err != nil {
			authToken = FetchUserReadTokenFromBrowser("")
		}

		userReadToken := config.TokenStructure{
			UserReadToken:          newToken.AccessToken,
			UserReadTokenExpiresIn: int64(newToken.ExpiresIn),
		}
		config.WriteTokenToHomeDirectory(&userReadToken, false)
		authToken = newToken.AccessToken
	}

	if len(authToken) == 0 {
		authToken = FetchUserReadTokenFromBrowser("")
	}
	return authToken
}
