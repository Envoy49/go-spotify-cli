package auth

import (
	"fmt"
	"github.com/envoy49/go-spotify-cli/commands/commandTypes"
	"github.com/pkg/browser"
)

func buildSpotifyURL(params *commandTypes.UrlParams) string {
	return fmt.Sprintf(
		"https://accounts.spotify.com/authorize?client_id=%s&response_type=code&redirect_uri=%s&scope=%s",
		params.ClientID,
		params.RedirectURI,
		params.RequestedScopes,
	)
}

func OpenAuthUrl(params *commandTypes.UrlParams) error {
	var authUrl = buildSpotifyURL(params)

	// Open URL in default browser
	err := browser.OpenURL(authUrl)
	if err != nil {
		return err
	}

	return nil
}
