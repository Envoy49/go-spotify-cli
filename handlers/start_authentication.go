package handlers

import (
	"fmt"
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
