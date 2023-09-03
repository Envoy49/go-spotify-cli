package handlers

import (
	"github.com/sirupsen/logrus"
	"go-spotify-cli/common"
	"go-spotify-cli/config"
	"go-spotify-cli/constants"
	"go-spotify-cli/utils"
)

func StartAuthTokenFlow(redirectionRoute string) {
	params := &common.UrlParams{
		ClientID:        config.GlobalConfig.ClientId,
		RedirectURI:     constants.ServerUrl + redirectionRoute,
		RequestedScopes: config.GlobalConfig.RequestedScopes,
	}

	if authUrlErr := utils.OpenAuthUrl(params); authUrlErr != nil {
		logrus.WithError(authUrlErr).Error("Error opening auth URL")
	}
}
