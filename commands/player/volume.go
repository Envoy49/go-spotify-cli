package player

import (
	"fmt"
	"net/url"

	"github.com/envoy49/go-spotify-cli/commands/cmdTypes"
	"github.com/envoy49/go-spotify-cli/config"

	"github.com/envoy49/go-spotify-cli/commands"
	"github.com/envoy49/go-spotify-cli/server"
	"github.com/sirupsen/logrus"
	"github.com/spf13/cobra"
)

var VolumeValue string

func volume(cfg *config.Config, accessToken string) {
	query := url.Values{}
	query.Set("volume_percent", VolumeValue)
	fullEndpoint := fmt.Sprintf("%s?%s", "/player/volume", query.Encode())
	params := &commands.PlayerParams{
		AccessToken: accessToken,
		Method:      "PUT",
		Endpoint:    spotifyPlayerEndpoint + fullEndpoint,
	}

	_, err := commands.Fetch(params)

	if err != nil {
		switch e := err.(type) {
		case cmdTypes.SpotifyAPIError:
			if e.Detail.Error.Message == "Player command failed: No active device found" {
				// Handle the case where no active device is found
				Device(cfg) // This function should ideally select or activate a default device
			}
		default:
			logrus.WithError(err).Error("Error setting volume")
			return
		}

	} else {
		logrus.Printf("Volume set to: %s%%", VolumeValue)
		Player(cfg)
	}
}

func VolumeCommand(cfg *config.Config) *cobra.Command {
	cmd := &cobra.Command{
		Use:   "volume [value]",
		Short: "Set volume",
		PreRunE: func(cmd *cobra.Command, args []string) error {
			if VolumeValue == "" {
				return fmt.Errorf("volume value is required")
			}
			return nil
		},
		Run: func(cmd *cobra.Command, args []string) {
			token := server.ReadUserModifyTokenOrFetchFromServer(cfg)
			volume(cfg, token)
		},
	}

	// Add a flag to the command
	cmd.Flags().StringVarP(&VolumeValue, "volume", "v", "", "Volume to set (0-100)")
	err := cmd.MarkFlagRequired("volume")
	if err != nil {
		logrus.WithError(err).Error("Error setting up volume command")
	}

	return cmd
}
