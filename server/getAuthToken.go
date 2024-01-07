package server

import (
	"github.com/envoy49/go-spotify-cli/auth"
	"github.com/envoy49/go-spotify-cli/config"
)

func ReadUserModifyTokenOrFetchFromServer() string {
	tokenInstance := config.ReadTokenFromHome(config.ModifyToken)
	if len(tokenInstance.ModifyToken.UserModifyRefreshToken) > 0 {
		newToken, err := auth.FetchAuthToken(&auth.FetchAuthTokenParams{
			RefreshToken: tokenInstance.ModifyToken.UserModifyRefreshToken,
		})

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
	tokenInstance := config.ReadTokenFromHome(config.ReadToken)
	if len(tokenInstance.ReadToken.UserReadRefreshToken) > 0 {
		newToken, err := auth.FetchAuthToken(&auth.FetchAuthTokenParams{
			RefreshToken: tokenInstance.ReadToken.UserReadRefreshToken,
		})

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

func ReadUserLibraryReadTokenOrFetchFromServer() string {
	tokenInstance := config.ReadTokenFromHome(config.LibraryRead)
	if len(tokenInstance.LibraryReadToken.UserLibraryReadRefreshToken) > 0 {
		newToken, err := auth.FetchAuthToken(&auth.FetchAuthTokenParams{
			RefreshToken: tokenInstance.LibraryReadToken.UserLibraryReadRefreshToken,
		})

		if err != nil {
			return FetchLibraryReadTokenFromBrowser()
		}

		userLibraryReadToken := config.CombinedTokenStructure{
			LibraryReadToken: config.UserLibraryReadTokenStructure{
				UserLibraryReadToken:          newToken.AccessToken,
				UserLibraryReadTokenExpiresIn: int64(newToken.ExpiresIn),
			},
		}
		config.WriteTokenToHomeDirectory(&userLibraryReadToken, false)
		return newToken.AccessToken
	}

	if len(tokenInstance.LibraryReadToken.UserLibraryReadToken) == 0 {
		return FetchLibraryReadTokenFromBrowser()
	}
	return tokenInstance.LibraryReadToken.UserLibraryReadToken
}

func FetchUserModifyTokenFromBrowser() string {
	config.GlobalConfig.RequestedScopes = config.UserModifyPlaybackStateScope
	cancel := StartServer(config.UserModifyPlaybackStateRoute)
	receivedToken := <-config.AuthTokenData
	cancel()
	return receivedToken.ModifyToken.UserModifyToken
}

func FetchUserReadTokenFromBrowser() string {
	config.GlobalConfig.RequestedScopes = config.UserReadPlaybackState
	cancel := StartServer(config.UserReadPlaybackStateRoute)
	receivedToken := <-config.AuthTokenData
	cancel()
	return receivedToken.ReadToken.UserReadToken
}

func FetchLibraryReadTokenFromBrowser() string {
	config.GlobalConfig.RequestedScopes = config.UserLibraryRead
	cancel := StartServer(config.UserLibraryReadRoute)
	receivedToken := <-config.AuthTokenData
	cancel()
	return receivedToken.LibraryReadToken.UserLibraryReadToken
}
