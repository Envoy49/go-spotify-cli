package handlers

import (
	"github.com/sirupsen/logrus"
	"go-spotify-cli/auth"
	"go-spotify-cli/config"
	"go-spotify-cli/constants"
	"net/http"
)

func UserModifyTokenHandler(w http.ResponseWriter, r *http.Request) {
	authCode := r.URL.Query().Get("code")

	accessToken, expiresIn, err := auth.FetchAuthToken(config.GlobalConfig.ClientId, config.GlobalConfig.ClientSecret, authCode, constants.ServerUrl+constants.AuthCallBackRoute)

	if err != nil {
		logrus.WithError(err).Error("Failed to get access token")
		return
	}

	userModifyTokenData := config.TokenStructure{
		UserModifyToken:          accessToken,
		UserModifyTokenExpiresIn: int64(expiresIn),
	}

	config.WriteTokenToHomeDirectory(&userModifyTokenData)
}
