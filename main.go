package main

import (
	"github.com/spf13/cobra"
	"go-spotify-cli/cmd/player/commands"
	"go-spotify-cli/config"
	"go-spotify-cli/constants"
	"go-spotify-cli/utils"
	"os"
)

func init() {
	config.LoadConfiguration()
}

func main() {
	var rootCmd = &cobra.Command{Use: constants.ProjectName}

	rootCmd.AddCommand(commands.PlayCommand, commands.PauseCommand, commands.NextCommand, commands.PreviousCommand, commands.DeviceCommand)
	if err := rootCmd.Execute(); err != nil {
		utils.PrintError("Error executing command", err)
		os.Exit(1)
	}
}
