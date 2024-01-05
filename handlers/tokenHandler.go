package handlers

import (
	"net/http"

	"github.com/envoy49/go-spotify-cli/auth"
	"github.com/envoy49/go-spotify-cli/config"
	"github.com/envoy49/go-spotify-cli/constants"
	"github.com/sirupsen/logrus"
)

func TokenHandler(w http.ResponseWriter, r *http.Request, tokenType constants.TokenType) {
	authCode := r.URL.Query().Get("code")
	callbackURL := constants.ServerUrl
	if tokenType == constants.ModifyToken {
		callbackURL += constants.UserModifyPlaybackStateRouteCallback
	} else if tokenType == constants.ReadToken {
		callbackURL += constants.UserReadPlaybackStateRouteCallback
	} else if tokenType == constants.LibraryRead {
		callbackURL += constants.UserLibraryReadRouteCallback
	}

	response, err := auth.FetchAuthToken(&auth.FetchAuthTokenParams{
		AuthCode:    authCode,
		RedirectURI: callbackURL,
	})

	if err != nil {
		logrus.WithError(err).Error("Failed to get access token")
		return
	}

	var tokenData config.CombinedTokenStructure

	if tokenType == constants.ModifyToken {
		tokenData = config.CombinedTokenStructure{
			ModifyToken: config.UserModifyTokenStructure{
				UserModifyToken:          response.AccessToken,
				UserModifyRefreshToken:   response.RefreshToken,
				UserModifyTokenExpiresIn: int64(response.ExpiresIn),
			},
		}
	}

	if tokenType == constants.ReadToken {
		tokenData = config.CombinedTokenStructure{
			ReadToken: config.UserReadTokenStructure{
				UserReadToken:          response.AccessToken,
				UserReadRefreshToken:   response.RefreshToken,
				UserReadTokenExpiresIn: int64(response.ExpiresIn),
			},
		}

	}

	if tokenType == constants.LibraryRead {
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
