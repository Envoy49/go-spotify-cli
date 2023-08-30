package commands

import (
	"fmt"
	"github.com/spf13/cobra"
	commands "go-spotify-cli/cmd/player"
	"go-spotify-cli/server"
	"go-spotify-cli/utils"
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

	_, _, err := commands.Player(params)

	if err != nil {
		utils.PrintError("Error setting volume", err)
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
		token := server.GetAuthTokenOrFetchFromServer()
		volume(token)
	},
}
