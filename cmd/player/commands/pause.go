package commands

import (
	"github.com/spf13/cobra"
	"go-spotify-cli/cmd/player"
	"go-spotify-cli/server"
	"go-spotify-cli/utils"
)

func pause(accessToken string) {
	params := &commands.PlayerParams{
		AccessToken: accessToken,
		Method:      "PUT",
		Endpoint:    "/pause",
	}
	_, _, err := commands.Player(params)

	if err != nil {
		utils.PrintError("Error pausing your track", err)
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
