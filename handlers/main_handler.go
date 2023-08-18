package handlers

import (
	"fmt"
	"go-spotify-cli/auth"
	"go-spotify-cli/cmd/player"
	"go-spotify-cli/common"
	"go-spotify-cli/constants"
	"go-spotify-cli/utils"
	"net/http"
)

func StartAuthentication(w http.ResponseWriter, r *http.Request) {
	params := &common.UrlParams{
		ClientID:        constants.ClientID,
		RedirectURI:     constants.ServerUrl + "/callback",
		RequestedScopes: constants.RequestedScopes,
	}

	if authUrlErr := utils.OpenAuthUrl(params); authUrlErr != nil {
		fmt.Println("Error opening auth URL", authUrlErr)
	}
}

func FetchAccessToken(w http.ResponseWriter, r *http.Request) {
	var token string

	authCode := r.URL.Query().Get("code")

	accessToken, expiresIn, err := auth.GetAccessToken(constants.ClientID, constants.ClientSecret, authCode, constants.ServerUrl+"/callback")
	if err != nil {
		utils.PrintError("Failed to get access token:", err)
		return
	}
	fmt.Println(constants.Green + "Expires in: " + fmt.Sprint(expiresIn) + " seconds" + constants.Reset)

	token = accessToken

	if err := utils.WriteJWTToken(token, expiresIn); err != nil {
		utils.PrintError("Failed to write JWT token:", err)
	}

	if playErr := player.Play(token); playErr != nil {
		utils.PrintError("Failed to get Play your track:", playErr)
	}

	fmt.Println(constants.Magenta + "Access token: " + fmt.Sprint(token) + " seconds" + constants.Reset)
}
