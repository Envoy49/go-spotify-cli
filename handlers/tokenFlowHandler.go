package handlers

import (
	"github.com/envoy49/go-spotify-cli/auth"
	"github.com/envoy49/go-spotify-cli/commands/commandTypes"
	"github.com/envoy49/go-spotify-cli/config"
	"github.com/sirupsen/logrus"
)

func StartAuthTokenFlow(cfg *config.Config, redirectionRoute string) {
	params := &commandTypes.UrlParams{
		ClientID:        cfg.ClientId,
		RedirectURI:     config.ServerUrl + redirectionRoute,
		RequestedScopes: cfg.RequestedScopes,
	}

	if authUrlErr := auth.OpenAuthUrl(params); authUrlErr != nil {
		logrus.WithError(authUrlErr).Error("Error opening auth URL")
		return
	}
}
