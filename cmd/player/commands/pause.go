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
	err := commands.Player(params)

	if err != nil {
		utils.PrintError("Error pausing your track", err)
	}
}

var PauseCommand = &cobra.Command{
	Use:   "pause",
	Short: "Pause spotify song",
	Run: func(cmd *cobra.Command, args []string) {
		token := utils.ReadJWTToken()
		if len(token) == 0 {
			server.StartAuthentication()
			receivedToken := <-utils.AuthToken
			server.InitiateShutdown()
			token = receivedToken
		}
		pause(token)
	},
}
