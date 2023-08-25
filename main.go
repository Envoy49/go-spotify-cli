package main

import (
	"github.com/spf13/cobra"
	"go-spotify-cli/cmd/player"
	"go-spotify-cli/config"
	"go-spotify-cli/constants"
	"go-spotify-cli/server"
	"go-spotify-cli/utils"
	"os"
)

func init() {
	config.LoadConfiguration()
}

func main() {
	var rootCmd = &cobra.Command{Use: constants.ProjectName}
	var cmdPlay = &cobra.Command{
		Use:   "play",
		Short: "Play spotify song",
		Run: func(cmd *cobra.Command, args []string) {
			token := utils.ReadJWTToken()
			if len(token) == 0 {
				server.StartAuthentication()
				receivedToken := <-utils.AuthToken
				server.InitiateShutdown()
				token = receivedToken
			}
			if playErr := player.Play(token); playErr != nil {
				utils.PrintError("Failed to get Play your track:", playErr)
			}
		},
	}

	rootCmd.AddCommand(cmdPlay)
	if err := rootCmd.Execute(); err != nil {
		utils.PrintError("Error executing command", err)
		os.Exit(1)
	}
}
