package handlers

import (
	"github.com/sirupsen/logrus"
	"go-spotify-cli/common"
	"go-spotify-cli/config"
	"go-spotify-cli/utils"
	"net/http"
)

func StartAuthTokenFlow(w http.ResponseWriter, r *http.Request, redirectionRoute string) {
	params := &common.UrlParams{
		ClientID:        config.GlobalConfig.ClientId,
		RedirectURI:     config.GlobalConfig.ServerUrl + redirectionRoute,
		RequestedScopes: config.GlobalConfig.RequestedScopes,
	}

	if authUrlErr := utils.OpenAuthUrl(params); authUrlErr != nil {
		logrus.WithError(authUrlErr).Error("Error opening auth URL")
	}
}
