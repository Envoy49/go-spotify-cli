package server

import (
	"github.com/envoy49/go-spotify-cli/auth"
	"github.com/envoy49/go-spotify-cli/config"
)

func ReadUserModifyTokenOrFetchFromServer(cfg *config.Config) string {
	tokenInstance := config.ReadTokenFromHome(config.ModifyToken)
	if len(tokenInstance.ModifyToken.UserModifyRefreshToken) > 0 {
		newToken, err := auth.FetchAuthToken(cfg, &auth.FetchAuthTokenParams{
			RefreshToken: tokenInstance.ModifyToken.UserModifyRefreshToken,
		})

		if err != nil {
			return FetchUserModifyTokenFromBrowser(cfg)
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
		return FetchUserModifyTokenFromBrowser(cfg)
	}

	return tokenInstance.ModifyToken.UserModifyToken
}

func ReadUserReadTokenOrFetchFromServer(cfg *config.Config) string {
	tokenInstance := config.ReadTokenFromHome(config.ReadToken)
	if len(tokenInstance.ReadToken.UserReadRefreshToken) > 0 {
		newToken, err := auth.FetchAuthToken(cfg, &auth.FetchAuthTokenParams{
			RefreshToken: tokenInstance.ReadToken.UserReadRefreshToken,
		})

		if err != nil {
			return FetchUserReadTokenFromBrowser(cfg)
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
		return FetchUserReadTokenFromBrowser(cfg)
	}
	return tokenInstance.ReadToken.UserReadToken
}

func ReadUserLibraryReadTokenOrFetchFromServer(cfg *config.Config) string {
	tokenInstance := config.ReadTokenFromHome(config.LibraryRead)
	if len(tokenInstance.LibraryReadToken.UserLibraryReadRefreshToken) > 0 {
		newToken, err := auth.FetchAuthToken(cfg, &auth.FetchAuthTokenParams{
			RefreshToken: tokenInstance.LibraryReadToken.UserLibraryReadRefreshToken,
		})

		if err != nil {
			return FetchLibraryReadTokenFromBrowser(cfg)
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
		return FetchLibraryReadTokenFromBrowser(cfg)
	}
	return tokenInstance.LibraryReadToken.UserLibraryReadToken
}

func FetchUserModifyTokenFromBrowser(cfg *config.Config) string {
	cfg.RequestedScopes = config.UserModifyPlaybackStateScope
	cancel := StartServer(cfg, config.UserModifyPlaybackStateRoute)
	receivedToken := <-config.AuthTokenData
	cancel()
	return receivedToken.ModifyToken.UserModifyToken
}

func FetchUserReadTokenFromBrowser(cfg *config.Config) string {
	cfg.RequestedScopes = config.UserReadPlaybackState
	cancel := StartServer(cfg, config.UserReadPlaybackStateRoute)
	receivedToken := <-config.AuthTokenData
	cancel()
	return receivedToken.ReadToken.UserReadToken
}

func FetchLibraryReadTokenFromBrowser(cfg *config.Config) string {
	cfg.RequestedScopes = config.UserLibraryRead
	cancel := StartServer(cfg, config.UserLibraryReadRoute)
	receivedToken := <-config.AuthTokenData
	cancel()
	return receivedToken.LibraryReadToken.UserLibraryReadToken
}
