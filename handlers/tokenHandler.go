package handlers

import (
	"net/http"

	"github.com/envoy49/go-spotify-cli/auth"
	"github.com/envoy49/go-spotify-cli/config"
	"github.com/sirupsen/logrus"
)

func TokenHandler(w http.ResponseWriter, r *http.Request, cfg *config.Config, tokenType config.TokenType) {
	authCode := r.URL.Query().Get("code")
	callbackURL := config.ServerUrl
	if tokenType == config.ModifyToken {
		callbackURL += config.UserModifyPlaybackStateRouteCallback
	} else if tokenType == config.ReadToken {
		callbackURL += config.UserReadPlaybackStateRouteCallback
	} else if tokenType == config.LibraryRead {
		callbackURL += config.UserLibraryReadRouteCallback
	}

	response, err := auth.FetchAuthToken(cfg, &auth.FetchAuthTokenParams{
		AuthCode:    authCode,
		RedirectURI: callbackURL,
	})

	if err != nil {
		logrus.WithError(err).Error("Failed to get access token")
		return
	}

	var tokenData config.CombinedTokenStructure

	if tokenType == config.ModifyToken {
		tokenData = config.CombinedTokenStructure{
			ModifyToken: config.UserModifyTokenStructure{
				UserModifyToken:          response.AccessToken,
				UserModifyRefreshToken:   response.RefreshToken,
				UserModifyTokenExpiresIn: int64(response.ExpiresIn),
			},
		}
	}

	if tokenType == config.ReadToken {
		tokenData = config.CombinedTokenStructure{
			ReadToken: config.UserReadTokenStructure{
				UserReadToken:          response.AccessToken,
				UserReadRefreshToken:   response.RefreshToken,
				UserReadTokenExpiresIn: int64(response.ExpiresIn),
			},
		}

	}

	if tokenType == config.LibraryRead {
		tokenData = config.CombinedTokenStructure{
			LibraryReadToken: config.UserLibraryReadTokenStructure{
				UserLibraryReadToken:          response.AccessToken,
				UserLibraryReadRefreshToken:   response.RefreshToken,
				UserLibraryReadTokenExpiresIn: int64(response.ExpiresIn),
			},
		}

	}

	config.WriteTokenToHomeDirectory(&tokenData, true)

	w.Header().Set("Content-Type", "text/html")
	if _, err := w.Write([]byte(content)); err != nil {
		logrus.Error(err)
	}
}
