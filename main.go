package main

import (
	"fmt"
	"go-spotify-cli/auth"
	"go-spotify-cli/common"
	"go-spotify-cli/constants"
	"go-spotify-cli/player"
	"go-spotify-cli/utils"
	"net/http"
)

func main() {
	params := &common.UrlParams{
		ClientID:        constants.ClientID,
		RedirectURI:     constants.ServerUrl,
		RequestedScopes: constants.RequestedScopes,
	}

	if authUrlErr := open_auth_url.OpenAuthUrl(params); authUrlErr != nil {
		panic(authUrlErr)
	}

	var token bool = false

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		authCode := r.URL.Query().Get("code")

		if token {
			return
		}

		accessToken, err := auth.GetAccessToken(constants.ClientID, constants.ClientSecret, authCode, constants.ServerUrl)
		if err != nil {
			fmt.Println("Failed to get access token:", err)
			return
		} else {
			token = true
		}
		if playErr := player.Play(accessToken); playErr != nil {
			panic(playErr)
		}
		fmt.Println("Access token:", accessToken)
	})

	fmt.Printf("Listening on %s\n", constants.ServerUrl)
	if err := http.ListenAndServe(constants.Port, nil); err != nil {
		panic(err)
	}

}
