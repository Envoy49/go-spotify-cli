package handlers

import (
	"github.com/envoy49/go-spotify-cli/auth"
	"github.com/envoy49/go-spotify-cli/commands/commandTypes"
	"github.com/envoy49/go-spotify-cli/config"
	"github.com/envoy49/go-spotify-cli/constants"
	"github.com/sirupsen/logrus"
)

func StartAuthTokenFlow(redirectionRoute string) {
	params := &commandTypes.UrlParams{
		ClientID:        config.GlobalConfig.ClientId,
		RedirectURI:     constants.ServerUrl + redirectionRoute,
		RequestedScopes: config.GlobalConfig.RequestedScopes,
	}

	if authUrlErr := auth.OpenAuthUrl(params); authUrlErr != nil {
		logrus.WithError(authUrlErr).Error("Error opening auth URL")
		return
	}
}
