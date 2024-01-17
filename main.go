package main

import (
	"os"

	"github.com/envoy49/go-spotify-cli/commands/player"
	"github.com/envoy49/go-spotify-cli/commands/search"

	"github.com/envoy49/go-spotify-cli/commands/flush"
	"github.com/envoy49/go-spotify-cli/config"
	"github.com/envoy49/go-spotify-cli/loader"
	"github.com/sirupsen/logrus"
	"github.com/spf13/cobra"
)

const (
	projectName = "go-spotify-cli"
	version     = "dev"
)

func main() {
	var cfg *config.Config
	loader.InitializeSpinner()

	cfgService := config.NewConfigService()
	cfg = cfgService.GetConfig()
	fetchType := cfgService.GetFetchType()

	var rootCmd = &cobra.Command{
		Use:     projectName,
		Version: version,
		PersistentPreRun: func(cmd *cobra.Command, args []string) {
			if cmd.Name() == "flush-tokens" || cmd.Name() == "flush-secrets" {
				return
			}
			loader.Start()
		},
		PersistentPostRun: func(cmd *cobra.Command, args []string) {
			loader.Stop()
		},
	}
	rootCmd.AddCommand(
		player.PlayCommand(cfg),
		player.PauseCommand(cfg),
		player.NextCommand(cfg),
		player.PreviousCommand(cfg),
		player.DeviceCommand(cfg),
		player.VolumeCommand(cfg),
		player.SavedCommand(cfg),
		search.SendSearchCommand(cfg),
		flush.FlushSecretsCommand(fetchType),
		flush.FlushTokensCommand,
	)

	if err := rootCmd.Execute(); err != nil {
		logrus.WithError(err).Error("Error executing command")
		os.Exit(1)
	}
}
