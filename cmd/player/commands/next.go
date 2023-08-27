package commands

import (
	"github.com/spf13/cobra"
	"go-spotify-cli/cmd/player"
	"go-spotify-cli/server"
	"go-spotify-cli/utils"
)

func next(accessToken string) {
	params := &commands.PlayerParams{
		AccessToken: accessToken,
		Method:      "POST",
		Endpoint:    "/next",
	}
	err := commands.Player(params)

	if err != nil {
		utils.PrintError("Error going to the next track", err)
	}
}

var NextCommand = &cobra.Command{
	Use:   "next",
	Short: "Next spotify song",
	Run: func(cmd *cobra.Command, args []string) {
		token := utils.ReadJWTToken()
		if len(token) == 0 {
			server.StartAuthentication()
			receivedToken := <-utils.AuthToken
			server.InitiateShutdown()
			token = receivedToken
		}
		next(token)
	},
}
