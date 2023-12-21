package handlers

import (
	"github.com/sirupsen/logrus"
	"go-spotify-cli/auth"
	"go-spotify-cli/config"
	"go-spotify-cli/constants"
	"go-spotify-cli/html"
	"net/http"
)

func TokenHandler(w http.ResponseWriter, r *http.Request, tokenType constants.TokenType) {
	authCode := r.URL.Query().Get("code")
	callbackURL := constants.ServerUrl
	if tokenType == constants.ModifyToken {
		callbackURL += constants.UserModifyPlaybackStateRouteCallback
	} else if tokenType == constants.ReadToken {
		callbackURL += constants.UserReadPlaybackStateRouteCallback
	}

	response, err := auth.FetchAuthToken(authCode, callbackURL)

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

	config.WriteTokenToHomeDirectory(&tokenData, true)

	w.Header().Set("Content-Type", "text/html")
	w.Write([]byte(html.ContentOfHTML))
}
