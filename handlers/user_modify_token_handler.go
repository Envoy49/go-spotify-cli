package handlers

import (
	"github.com/sirupsen/logrus"
	"go-spotify-cli/auth"
	"go-spotify-cli/config"
	"go-spotify-cli/constants"
	"go-spotify-cli/html"
	"net/http"
)

func UserModifyTokenHandler(w http.ResponseWriter, r *http.Request) {
	authCode := r.URL.Query().Get("code")

	response, err := auth.FetchAuthToken(
		authCode,
		constants.ServerUrl+constants.UserModifyPlaybackStateRouteCallback,
	)

	if err != nil {
		logrus.WithError(err).Error("Failed to get access token")
		return
	}

	userModifyTokenData := config.TokenStructure{
		UserModifyToken:          response.AccessToken,
		UserModifyRefreshToken:   response.RefreshToken,
		UserModifyTokenExpiresIn: int64(response.ExpiresIn),
	}

	config.WriteTokenToHomeDirectory(&userModifyTokenData, true)

	w.Header().Set("Content-Type", "text/html")
	w.Write([]byte(html.ContentOfHTML))
}
