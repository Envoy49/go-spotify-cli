package handlers

import (
	"fmt"
	"go-spotify-cli/auth"
	"go-spotify-cli/constants"
	"go-spotify-cli/player"
	"go-spotify-cli/utils"
	"net/http"
)

func FetchAccessToken(w http.ResponseWriter, r *http.Request) {
	var token string

	token, cacheErr := utils.ReadJWTToken()
	if cacheErr != nil {
		fmt.Println("Failed to get access token from cache:", cacheErr)
		return
	}
	if len(token) == 0 {
		authCode := r.URL.Query().Get("code")
		accessToken, expiresIn, err := auth.GetAccessToken(constants.ClientID, constants.ClientSecret, authCode, constants.ServerUrl)
		if err != nil {
			fmt.Println("Failed to get access token:", err)
			return
		}
		fmt.Println("===========> expiresIn", expiresIn)
		token = accessToken

		if err := utils.WriteJWTToken(token); err != nil {
			fmt.Println("Failed to write JWT token:", err)
		}
	}

	if playErr := player.Play(token); playErr != nil {
		fmt.Println("Failed to get Play your track:", playErr)
	}

	fmt.Println("Access token:", token)
}
