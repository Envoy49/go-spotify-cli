package commands

import (
	"github.com/sirupsen/logrus"
	"github.com/spf13/cobra"
	"go-spotify-cli/cmd/player"
	"go-spotify-cli/common"
	"go-spotify-cli/server"
)

func next(accessToken string) {
	params := &commands.PlayerParams{
		AccessToken: accessToken,
		Method:      "POST",
		Endpoint:    "/next",
	}
	_, err := commands.FetchCommand(params)

	if err != nil {
		switch e := err.(type) {
		case common.SpotifyAPIError:
			if e.Detail.Error.Message == "Player command failed: No active device found" {
				Device()
			}
		}

		logrus.WithError(err).Error("Error going to the next track")
	} else {
		logrus.Info("Playing next track")
		Player()
	}
}

var NextCommand = &cobra.Command{
	Use:   "next",
	Short: "Next spotify song",
	Run: func(cmd *cobra.Command, args []string) {
		token := server.ReadUserModifyTokenOrFetchFromServer()
		next(token)
	},
}
