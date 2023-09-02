package commands

import (
	"github.com/sirupsen/logrus"
	"github.com/spf13/cobra"
	"go-spotify-cli/cmd/player"
	"go-spotify-cli/server"
)

func pause(accessToken string) {
	params := &commands.PlayerParams{
		AccessToken: accessToken,
		Method:      "PUT",
		Endpoint:    "/pause",
	}
	_, _, err := commands.Player(params)

	if err != nil {
		logrus.WithError(err).Error("Error pausing your track")
	} else {
		logrus.Info("Paused")
	}
}

var PauseCommand = &cobra.Command{
	Use:   "pause",
	Short: "Pause spotify song",
	Run: func(cmd *cobra.Command, args []string) {
		token := server.GetAuthTokenOrFetchFromServer()
		pause(token)
	},
}
