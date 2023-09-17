package commands

import (
	"fmt"
	"github.com/sirupsen/logrus"
	"github.com/spf13/cobra"
	"go-spotify-cli/cmd/player"
	"go-spotify-cli/common"
	"go-spotify-cli/server"
	"net/url"
)

var VolumeValue string

func volume(accessToken string) {
	query := url.Values{}
	query.Set("volume_percent", VolumeValue)
	fullEndpoint := fmt.Sprintf("%s?%s", "/volume", query.Encode())
	params := &commands.PlayerParams{
		AccessToken: accessToken,
		Method:      "PUT",
		Endpoint:    fullEndpoint,
	}

	_, err := commands.FetchCommand(params)

	if err != nil {
		switch e := err.(type) {
		case common.SpotifyAPIError:
			if e.Detail.Error.Message == "Player command failed: No active device found" {
				Device()
			}
		}

		logrus.WithError(err).Error("Error setting volume")

	} else {
		logrus.Info("Changed volume")
		Player()
	}
}

var VolumeCommand = &cobra.Command{
	Use:   "volume",
	Short: "Set volume",
	PreRunE: func(cmd *cobra.Command, args []string) error {
		if VolumeValue == "" {
			return fmt.Errorf("volume value is required")
		}
		return nil
	},
	Run: func(cmd *cobra.Command, args []string) {
		token := server.ReadUserModifyTokenOrFetchFromServer()
		volume(token)
	},
}
