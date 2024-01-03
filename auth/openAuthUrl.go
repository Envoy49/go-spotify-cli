package auth

import (
	"fmt"
	"github.com/pkg/browser"
	"go-spotify-cli/types"
)

func buildSpotifyURL(params *types.UrlParams) string {
	return fmt.Sprintf(
		"https://accounts.spotify.com/authorize?client_id=%s&response_type=code&redirect_uri=%s&scope=%s",
		params.ClientID,
		params.RedirectURI,
		params.RequestedScopes,
	)
}

func OpenAuthUrl(params *types.UrlParams) error {
	var authUrl = buildSpotifyURL(params)

	// Open URL in default browser
	err := browser.OpenURL(authUrl)
	if err != nil {
		return err
	}

	return nil
}
