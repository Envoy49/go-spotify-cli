package player

import (
	"github.com/envoy49/go-spotify-cli/commands"
	"github.com/envoy49/go-spotify-cli/server"
	"github.com/envoy49/go-spotify-cli/types"
	"github.com/sirupsen/logrus"
	"github.com/spf13/cobra"
)

func previous(accessToken string) {
	params := &commands.PlayerParams{
		AccessToken: accessToken,
		Method:      "POST",
		Endpoint:    spotifyPlayerEndpoint + "/player/previous",
	}
	_, err := commands.Fetch(params)

	if err != nil {
		switch e := err.(type) {
		case types.SpotifyAPIError:
			if e.Detail.Error.Message == "Player command failed: No active device found" {
				Device()
			}
		default:
			logrus.WithError(err).Error("Error going to the previous track")
			return
		}
	} else {
		Player()
	}
}

var PreviousCommand = &cobra.Command{
	Use:   "previous",
	Short: "Previous spotify song",
	Run: func(cmd *cobra.Command, args []string) {
		token := server.ReadUserModifyTokenOrFetchFromServer()
		previous(token)
	},
}
