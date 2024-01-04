package main

import (
	"os"

	"github.com/envoy49/go-spotify-cli/commands/flush"
	"github.com/envoy49/go-spotify-cli/commands/player"
	"github.com/envoy49/go-spotify-cli/commands/search"
	"github.com/envoy49/go-spotify-cli/config"
	"github.com/envoy49/go-spotify-cli/loader"
	"github.com/envoy49/go-spotify-cli/prompt"
	"github.com/sirupsen/logrus"
	"github.com/spf13/cobra"
)

const (
	projectName = "go-spotify-cli"
)

func init() {
	config.LoadConfiguration()
}

var Version string

func main() {
	loader.InitializeSpinner()

	var rootCmd = &cobra.Command{
		Use:     projectName,
		Version: Version,
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

	if err := rootCmd.Execute(); err != nil {
		logrus.WithError(err).Error("Error executing command")
		os.Exit(1)
	}
}
