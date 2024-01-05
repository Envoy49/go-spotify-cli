package player

import (
	"github.com/envoy49/go-spotify-cli/commands"
	"github.com/envoy49/go-spotify-cli/server"
	"github.com/envoy49/go-spotify-cli/types"
	"github.com/sirupsen/logrus"
	"github.com/spf13/cobra"
)

func pause(accessToken string) {
	params := &commands.PlayerParams{
		AccessToken: accessToken,
		Method:      "PUT",
		Endpoint:    spotifyPlayerEndpoint + "/player/pause",
	}
	_, err := commands.Fetch(params)

	if err != nil {
		switch e := err.(type) {
		case types.SpotifyAPIError:
			if e.Detail.Error.Message == "Player command failed: No active device found" {
				Device()
			}
		default:
			logrus.WithError(err).Error("Error pausing your track")
			return
		}

	} else {
		logrus.Println("Paused")
	}
}

var PauseCommand = &cobra.Command{
	Use:   "pause",
	Short: "Pause spotify song",
	Run: func(cmd *cobra.Command, args []string) {
		token := server.ReadUserModifyTokenOrFetchFromServer()
		pause(token)
	},
}
