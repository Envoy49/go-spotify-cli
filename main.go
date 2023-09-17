package main

import (
	"github.com/sirupsen/logrus"
	"github.com/spf13/cobra"
	"go-spotify-cli/cmd/player/commands"
	"go-spotify-cli/config"
	"go-spotify-cli/constants"
	"os"
)

func init() {
	config.LoadConfiguration()
}

func main() {
	var rootCmd = &cobra.Command{
		Use:              constants.ProjectName,
		PersistentPreRun: config.EnvVarsPrompt,
	}

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
		logrus.WithError(err).Error("Error running volume command")
	}
	if err := rootCmd.Execute(); err != nil {
		logrus.WithError(err).Error("Error executing command")
		os.Exit(1)
	}
}
