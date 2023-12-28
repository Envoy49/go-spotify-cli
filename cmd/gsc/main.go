package main

import (
	"github.com/sirupsen/logrus"
	"github.com/spf13/cobra"
	"go-spotify-cli/commands/flush"
	"go-spotify-cli/commands/player"
	"go-spotify-cli/commands/search"
	"go-spotify-cli/config"
	"go-spotify-cli/constants"
	"go-spotify-cli/loader"
	"go-spotify-cli/prompt"
	"os"
)

func init() {
	config.LoadConfiguration()
}

func main() {
	loader.InitializeSpinner()

	var rootCmd = &cobra.Command{
		Use: constants.ProjectName,
		PersistentPreRun: func(cmd *cobra.Command, args []string) {
			if cmd.Name() == "flush-tokens" || cmd.Name() == "flush-secrets" {
				return
			}
			loader.Start()
			prompt.EnvVarsPrompt()
		},
		PersistentPostRun: func(cmd *cobra.Command, args []string) {
			loader.Stop()
		},
	}

	rootCmd.AddCommand(
		player.PlayCommand,
		player.PauseCommand,
		player.NextCommand,
		player.PreviousCommand,
		player.DeviceCommand,
		player.VolumeCommand,
		player.SavedCommand,
		search.SendSearchCommand,
		flush.FlushTokensCommand,
		flush.FlushSecretsCommand,
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
