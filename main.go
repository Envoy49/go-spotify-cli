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

	rootCmd.AddCommand(
		commands.PlayCommand,
		commands.PauseCommand,
		commands.NextCommand,
		commands.PreviousCommand,
		commands.DeviceCommand,
		commands.VolumeCommand,
	)

	commands.VolumeCommand.Flags().StringVarP(&commands.VolumeValue, "volume", "v", "", "Volume to add")
	err := commands.VolumeCommand.MarkFlagRequired("volume")
	if err != nil {
		utils.PrintError("Error running volume command", err)
	}

	if err := rootCmd.Execute(); err != nil {
		utils.PrintError("Error executing command", err)
		os.Exit(1)
	}
}
