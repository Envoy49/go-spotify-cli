package commands

import (
	"github.com/spf13/cobra"
	"go-spotify-cli/cmd/player"
	"go-spotify-cli/server"
	"go-spotify-cli/utils"
)

func play(accessToken string) {
	params := &commands.PlayerParams{
		AccessToken: accessToken,
		Method:      "PUT",
		Endpoint:    "/play",
	}
	_, _, err := commands.Player(params)

	if err != nil {
		utils.PrintError("Error playing your track", err)
	}
}

var PlayCommand = &cobra.Command{
	Use:   "play",
	Short: "Play spotify song",
	Run: func(cmd *cobra.Command, args []string) {
		token := server.GetAuthTokenOrFetchFromServer()
		play(token)
	},
}
