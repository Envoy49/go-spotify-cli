package player

import (
	"github.com/sirupsen/logrus"
	"github.com/spf13/cobra"
	"go-spotify-cli/commands"
	"go-spotify-cli/constants"
	"go-spotify-cli/server"
	"go-spotify-cli/types"
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
