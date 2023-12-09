package handlers

import (
	"github.com/sirupsen/logrus"
	"go-spotify-cli/auth"
	"go-spotify-cli/config"
	"go-spotify-cli/constants"
	"go-spotify-cli/html"
	"net/http"
)

func UserReadTokenHandler(w http.ResponseWriter, r *http.Request) {
	authCode := r.URL.Query().Get("code")

	response, err := auth.FetchAuthToken(
		authCode,
		constants.ServerUrl+constants.UserReadPlaybackStateRouteCallback,
	)

	if err != nil {
		logrus.WithError(err).Error("Failed to get access token")
		return
	}

	userReadTokenData := config.TokenStructure{
		UserReadToken:          response.AccessToken,
		UserReadRefreshToken:   response.RefreshToken,
		UserReadTokenExpiresIn: int64(response.ExpiresIn),
	}

	config.WriteTokenToHomeDirectory(&userReadTokenData, true)

	w.Header().Set("Content-Type", "text/html")
	w.Write([]byte(html.ContentOfHTML))
}
