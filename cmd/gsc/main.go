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

	var silentMode bool
	var silentProgressMode bool

	var rootCmd = &cobra.Command{
		Use: constants.ProjectName,
		PersistentPreRun: func(cmd *cobra.Command, args []string) {
			if silentMode || silentProgressMode {
				return
			}

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

	rootCmd.PersistentFlags().BoolVarP(&silentMode, "silent", "s", false, "Run in silent mode without user interaction")
	rootCmd.PersistentFlags().BoolVarP(&silentProgressMode, "silent-progress", "p", false, "Run in silent mode but show progress")

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

	if err := rootCmd.Execute(); err != nil {
		logrus.WithError(err).Error("Error executing command")
		os.Exit(1)
	}
}
