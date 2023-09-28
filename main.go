package main

import (
	"github.com/sirupsen/logrus"
	"github.com/spf13/cobra"
	"go-spotify-cli/cmd/player"
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
		player.PlayCommand,
		player.PauseCommand,
		player.NextCommand,
		player.PreviousCommand,
		player.DeviceCommand,
		player.VolumeCommand,
	)

	player.VolumeCommand.Flags().StringVarP(&player.VolumeValue, "volume", "v", "", "Volume to add")
	err := player.VolumeCommand.MarkFlagRequired("volume")
	if err != nil {
		logrus.WithError(err).Error("Error running volume command")
	}
	if err := rootCmd.Execute(); err != nil {
		logrus.WithError(err).Error("Error executing command")
		os.Exit(1)
	}
}
