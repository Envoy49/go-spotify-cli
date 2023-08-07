package utils

import (
	"fmt"
	"go-spotify-cli/common"
	"os/exec"
	"runtime"
)

func buildSpotifyURL(params *common.UrlParams) string {
	return fmt.Sprintf(
		"https://accounts.spotify.com/authorize?client_id=%s&response_type=code&redirect_uri=%s&scope=%s",
		params.ClientID,
		params.RedirectURI,
		params.RequestedScopes,
	)
}

func OpenAuthUrl(params *common.UrlParams) error {
	var authUrl = buildSpotifyURL(params)
	var cmd string
	var args []string

	switch runtime.GOOS {
	case "windows":
		cmd = "cmd"
		args = []string{"/c", "start"}
	case "darwin":
		cmd = "open"
	default: // "linux", "freebsd", "openbsd", "netbsd"
		cmd = "xdg-open"
	}
	args = append(args, authUrl)
	return exec.Command(cmd, args...).Start()
}
