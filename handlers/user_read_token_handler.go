package handlers

import (
	"github.com/sirupsen/logrus"
	"go-spotify-cli/auth"
	"go-spotify-cli/config"
	"go-spotify-cli/constants"
	"net/http"
)

func UserReadTokenHandler(w http.ResponseWriter, r *http.Request) {
	authCode := r.URL.Query().Get("code")

	accessToken, expiresIn, err := auth.FetchAuthToken(
		config.GlobalConfig.ClientId,
		config.GlobalConfig.ClientSecret,
		authCode,
		constants.ServerUrl+constants.UserReadPlaybackStateRouteCallback,
	)

	if err != nil {
		logrus.WithError(err).Error("Failed to get access token")
		return
	}

	userReadTokenData := config.TokenStructure{
		UserReadToken:          accessToken,
		UserReadTokenExpiresIn: int64(expiresIn),
	}

	config.WriteTokenToHomeDirectory(&userReadTokenData)
}
