package handlers

import (
	"fmt"
	"github.com/sirupsen/logrus"
	"go-spotify-cli/auth"
	"go-spotify-cli/config"
	"go-spotify-cli/constants"
	"go-spotify-cli/utils"
	"net/http"
)

func FetchAccessToken(w http.ResponseWriter, r *http.Request) {
	var token string
	authCode := r.URL.Query().Get("code")

	accessToken, expiresIn, err := auth.GetAccessToken(config.GlobalConfig.ClientId, config.GlobalConfig.ClientSecret, authCode, constants.ServerUrl+constants.AuthCallBackRoute)
	if err != nil {
		logrus.WithError(err).Error("Failed to get access token")
		return
	}
	logrus.Info("Token expires in: " + fmt.Sprint(expiresIn) + " seconds")

	token = accessToken

	if err := utils.WriteJWTToken(token, expiresIn); err != nil {
		logrus.WithError(err).Error("Failed to write JWT token")
	}
}
