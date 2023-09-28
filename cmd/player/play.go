package player

import (
	"github.com/sirupsen/logrus"
	"github.com/spf13/cobra"
	"go-spotify-cli/cmd"
	"go-spotify-cli/common"
	"go-spotify-cli/constants"
	"go-spotify-cli/server"
)

func play(accessToken string) {
	params := &cmd.PlayerParams{
		AccessToken: accessToken,
		Method:      "PUT",
		Endpoint:    constants.SpotifyPlayerEndpoint + "/player/play",
	}
	_, err := cmd.FetchCommand(params)

	if err != nil {
		switch e := err.(type) {
		case common.SpotifyAPIError:
			if e.Detail.Error.Message == "Player command failed: No active device found" {
				Device()
			} else {
				// Some other action for other SpotifyAPIError types.
				logrus.WithError(err).Error("Some other SpotifyAPIError occurred")
			}
		default:
			logrus.WithError(err).Error("Error playing your track")
		}
	} else {
		logrus.Info("Playing")
		Player()
	}
}

var PlayCommand = &cobra.Command{
	Use:   "play",
	Short: "Play spotify song",
	Run: func(cmd *cobra.Command, args []string) {
		token := server.ReadUserModifyTokenOrFetchFromServer()
		play(token)
	},
}
