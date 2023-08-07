package main

import (
	"fmt"
	"go-spotify-cli/common"
	"go-spotify-cli/constants"
	"go-spotify-cli/handlers"
	"go-spotify-cli/utils"
	"net/http"
)

func main() {
	params := &common.UrlParams{
		ClientID:        constants.ClientID,
		RedirectURI:     constants.ServerUrl,
		RequestedScopes: constants.RequestedScopes,
	}

	if authUrlErr := utils.OpenAuthUrl(params); authUrlErr != nil {
		fmt.Println("Error opening auth URL", authUrlErr)
	}

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		fmt.Println(r.Method)
		handlers.FetchAccessToken(w, r)
	})

	fmt.Printf("Listening on %s\n", constants.ServerUrl)
	if err := http.ListenAndServe(constants.Port, nil); err != nil {
		fmt.Println("Error listening starting the server", err)
	}

}
