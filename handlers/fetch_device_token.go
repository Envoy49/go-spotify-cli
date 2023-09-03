package handlers

import (
	"fmt"
	"github.com/sirupsen/logrus"
	"go-spotify-cli/auth"
	"go-spotify-cli/config"
	"go-spotify-cli/constants"
	"net/http"
)

var DeviceToken = make(chan string)

func FetchDeviceToken(w http.ResponseWriter, r *http.Request) {
	authCode := r.URL.Query().Get("code")

	accessToken, expiresIn, err := auth.GetAccessToken(config.GlobalConfig.ClientId, config.GlobalConfig.ClientSecret, authCode, constants.ServerUrl+constants.DeviceCallBackRoute)
	if err != nil {
		logrus.WithError(err).Error("Failed to get access token")
		return
	}

	logrus.Info("Token expires in: " + fmt.Sprint(expiresIn) + " seconds")

	DeviceToken <- accessToken

}
