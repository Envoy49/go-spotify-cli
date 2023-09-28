package player

import (
	"github.com/sirupsen/logrus"
	"github.com/spf13/cobra"
	"go-spotify-cli/cmd"
	"go-spotify-cli/common"
	"go-spotify-cli/constants"
	"go-spotify-cli/server"
)

func previous(accessToken string) {
	params := &cmd.PlayerParams{
		AccessToken: accessToken,
		Method:      "POST",
		Endpoint:    constants.SpotifyPlayerEndpoint + "/player/previous",
	}
	_, err := cmd.FetchCommand(params)

	if err != nil {
		switch e := err.(type) {
		case common.SpotifyAPIError:
			if e.Detail.Error.Message == "Player command failed: No active device found" {
				Device()
			}
		}

		logrus.WithError(err).Error("Error going to the previous track")
	} else {
		logrus.Info("Changed volume")
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
