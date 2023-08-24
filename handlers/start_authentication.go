package handlers

import (
	"go-spotify-cli/common"
	"go-spotify-cli/config"
	"go-spotify-cli/utils"
	"net/http"
)

func StartAuthentication(w http.ResponseWriter, r *http.Request) {
	params := &common.UrlParams{
		ClientID:        config.GlobalConfig.ClientId,
		RedirectURI:     config.GlobalConfig.ServerUrl + "/callback",
		RequestedScopes: config.GlobalConfig.RequestedScopes,
	}

	if authUrlErr := utils.OpenAuthUrl(params); authUrlErr != nil {
		utils.PrintError("Error opening auth URL", authUrlErr)
	}
}
