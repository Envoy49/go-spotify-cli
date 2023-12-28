package handlers

import (
	"github.com/sirupsen/logrus"
	"go-spotify-cli/auth"
	"go-spotify-cli/config"
	"go-spotify-cli/constants"
	"go-spotify-cli/types"
)

func StartAuthTokenFlow(redirectionRoute string) {
	params := &types.UrlParams{
		ClientID:        config.GlobalConfig.ClientId,
		RedirectURI:     constants.ServerUrl + redirectionRoute,
		RequestedScopes: config.GlobalConfig.RequestedScopes,
	}

	if authUrlErr := auth.OpenAuthUrl(params); authUrlErr != nil {
		logrus.WithError(authUrlErr).Error("Error opening auth URL")
		return
	}
}
