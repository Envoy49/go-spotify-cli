package main

import (
	"github.com/briandowns/spinner"
	"github.com/sirupsen/logrus"
	"github.com/spf13/cobra"
	"go-spotify-cli/commands/player"
	"go-spotify-cli/commands/search"
	"go-spotify-cli/config"
	"go-spotify-cli/constants"
	"os"
	"time"
)

func init() {
	config.LoadConfiguration()
}

func main() {
	s := spinner.New(spinner.CharSets[25], 50*time.Millisecond)
	s.Color("bold", "blue")

	var rootCmd = &cobra.Command{
		Use: constants.ProjectName,
		PersistentPreRun: func(cmd *cobra.Command, args []string) {
			if cmd.Name() != "search" {
				s.Start()
			}
			config.EnvVarsPrompt()
		},
		PersistentPostRun: func(cmd *cobra.Command, args []string) {
			if cmd.Name() != "search" {
				s.Stop()
			}
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
