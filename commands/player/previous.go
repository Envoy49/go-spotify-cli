package player

import (
	"github.com/sirupsen/logrus"
	"github.com/spf13/cobra"
	"go-spotify-cli/commands"
	"go-spotify-cli/constants"
	"go-spotify-cli/server"
	"go-spotify-cli/types"
)

func previous(accessToken string) {
	params := &commands.PlayerParams{
		AccessToken: accessToken,
		Method:      "POST",
		Endpoint:    constants.SpotifyPlayerEndpoint + "/player/previous",
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
