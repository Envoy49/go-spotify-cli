package server

import (
	"go-spotify-cli/auth"
	"go-spotify-cli/config"
	"go-spotify-cli/constants"
	"go-spotify-cli/types"
)

func ReadUserModifyTokenOrFetchFromServer() string {
	tokenInstance := config.ReadTokenFromHome(constants.ModifyToken)
	if len(tokenInstance.ModifyToken.UserModifyRefreshToken) > 0 {
		newToken, err := auth.FetchAuthToken(&types.FetchAuthTokenParams{
			RefreshToken: tokenInstance.ModifyToken.UserModifyRefreshToken,
		})

		if err != nil {
			return FetchUserModifyTokenFromBrowser()
		}

		userModifyToken := types.CombinedTokenStructure{
			ModifyToken: types.UserModifyTokenStructure{
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
		newToken, err := auth.FetchAuthToken(&types.FetchAuthTokenParams{
			RefreshToken: tokenInstance.ReadToken.UserReadRefreshToken,
		})

		if err != nil {
			return FetchUserReadTokenFromBrowser()
		}

		userReadToken := types.CombinedTokenStructure{
			ReadToken: types.UserReadTokenStructure{
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
	cancel := StartServer(constants.UserModifyPlaybackStateRoute)
	receivedToken := <-config.AuthTokenData
	cancel()
	return receivedToken.ModifyToken.UserModifyToken
}

func FetchUserReadTokenFromBrowser() string {
	config.GlobalConfig.RequestedScopes = constants.UserReadPlaybackState
	cancel := StartServer(constants.UserReadPlaybackStateRoute)
	receivedToken := <-config.AuthTokenData
	cancel()
	return receivedToken.ReadToken.UserReadToken
}
