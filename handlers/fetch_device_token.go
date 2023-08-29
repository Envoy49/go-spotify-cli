package handlers

import (
	"fmt"
	"go-spotify-cli/auth"
	"go-spotify-cli/config"
	"go-spotify-cli/constants"
	"go-spotify-cli/utils"
	"net/http"
)

var DeviceToken = make(chan string)

func FetchDeviceToken(w http.ResponseWriter, r *http.Request) {
	authCode := r.URL.Query().Get("code")

	accessToken, expiresIn, err := auth.GetAccessToken(config.GlobalConfig.ClientId, config.GlobalConfig.ClientSecret, authCode, config.GlobalConfig.ServerUrl+constants.DeviceCallBackRoute)
	if err != nil {
		utils.PrintError("Failed to get access token:", err)
		return
	}
	fmt.Println(constants.Green + "Token expires in: " + fmt.Sprint(expiresIn) + " seconds" + constants.Reset)

	DeviceToken <- accessToken

}
