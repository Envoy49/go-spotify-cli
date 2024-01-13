package main

import (
	"github.com/envoy49/go-spotify-cli/commands/player"
	"github.com/envoy49/go-spotify-cli/commands/search"
	"os"

	"github.com/envoy49/go-spotify-cli/commands/flush"
	"github.com/envoy49/go-spotify-cli/config"
	"github.com/envoy49/go-spotify-cli/loader"
	"github.com/sirupsen/logrus"
	"github.com/spf13/cobra"
)

const (
	projectName = "go-spotify-cli"
)

var Version string = "v1.0.65" // hardcode version for now until issue with dynamic assignment is resolved

func main() {
	var cfg *config.Config
	loader.InitializeSpinner()

	cfgService := config.NewConfigService()
	cfg = cfgService.GetConfig()
	fetchType := cfgService.GetFetchType()

	var rootCmd = &cobra.Command{
		Use:     projectName,
		Version: Version,
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
