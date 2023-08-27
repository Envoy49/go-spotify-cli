package commands

import (
	"github.com/spf13/cobra"
	"go-spotify-cli/cmd/player"
	"go-spotify-cli/server"
	"go-spotify-cli/utils"
)

func previous(accessToken string) {
	params := &commands.PlayerParams{
		AccessToken: accessToken,
		Method:      "POST",
		Endpoint:    "/previous",
	}
	err := commands.Player(params)

	if err != nil {
		utils.PrintError("Error going to the previous track", err)
	}
}

var PreviousCommand = &cobra.Command{
	Use:   "previous",
	Short: "Previous spotify song",
	Run: func(cmd *cobra.Command, args []string) {
		token := utils.ReadJWTToken()
		if len(token) == 0 {
			server.StartAuthentication()
			receivedToken := <-utils.AuthToken
			server.InitiateShutdown()
			token = receivedToken
		}
		previous(token)
	},
}
