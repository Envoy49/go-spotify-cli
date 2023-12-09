package main

import (
	"github.com/sirupsen/logrus"
	"github.com/spf13/cobra"
	"go-spotify-cli/commands"
	"go-spotify-cli/commands/player"
	"go-spotify-cli/commands/search"
	"go-spotify-cli/config"
	"go-spotify-cli/constants"
	"go-spotify-cli/spinnerInstance"
	"os"
)

func init() {
	config.LoadConfiguration()
}

func main() {
	spinnerInstance.InitializeSpinner()

	var rootCmd = &cobra.Command{
		Use: constants.ProjectName,
		PersistentPreRun: func(cmd *cobra.Command, args []string) {
			spinnerInstance.Start()
			config.EnvVarsPrompt()
		},
		PersistentPostRun: func(cmd *cobra.Command, args []string) {
			spinnerInstance.Stop()
		},
	}

	rootCmd.AddCommand(
		player.PlayCommand,
		player.PauseCommand,
		player.NextCommand,
		player.PreviousCommand,
		player.DeviceCommand,
		player.VolumeCommand,
		search.SendSearchCommand,
		commands.FlushAllTokensCommand,
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
