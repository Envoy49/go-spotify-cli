package player

import (
	"github.com/envoy49/go-spotify-cli/commands"
	"github.com/envoy49/go-spotify-cli/constants"
	"github.com/envoy49/go-spotify-cli/server"
	"github.com/envoy49/go-spotify-cli/types"
	"github.com/sirupsen/logrus"
	"github.com/spf13/cobra"
)

func Next(accessToken string, callPlayer bool) {
	params := &commands.PlayerParams{
		AccessToken: accessToken,
		Method:      "POST",
		Endpoint:    constants.SpotifyPlayerEndpoint + "/player/next",
	}
	_, err := commands.Fetch(params)
	if err != nil {
		switch e := err.(type) {
		case types.SpotifyAPIError:
			if e.Detail.Error.Message == "Player command failed: No active device found" {
				Device()
				return
			}
		default:
			logrus.WithError(err).Error("Error going to the next track")
			return
		}
	} else {
		if callPlayer {
			Player()
		}
	}
}

var NextCommand = &cobra.Command{
	Use:   "next",
	Short: "Next spotify song",
	Run: func(cmd *cobra.Command, args []string) {
		token := server.ReadUserModifyTokenOrFetchFromServer()
		Next(token, true)
	},
}
