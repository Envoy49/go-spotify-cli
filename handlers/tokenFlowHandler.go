package handlers

import (
	"github.com/envoy49/go-spotify-cli/auth"
	"github.com/envoy49/go-spotify-cli/config"
	"github.com/envoy49/go-spotify-cli/constants"
	"github.com/envoy49/go-spotify-cli/types"
	"github.com/sirupsen/logrus"
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
