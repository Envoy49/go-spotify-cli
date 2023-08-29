package commands

import (
	"fmt"
	"github.com/spf13/cobra"
	"go-spotify-cli/cmd/player"
	"go-spotify-cli/config"
	"go-spotify-cli/server"
)

func device() {
	config.GlobalConfig.RequestedScopes = "user-read-playback-state"
	token := server.FetchDeviceTokenFromBrowser()
	params := &commands.PlayerParams{
		AccessToken: token,
		Method:      "GET",
		Endpoint:    "/devices",
	}
	_, response, _ := commands.Player(params)
	fmt.Println("List of available devices:", response)
}

var DeviceCommand = &cobra.Command{
	Use:   "device",
	Short: "Get all connected devices",
	Run: func(cmd *cobra.Command, args []string) {
		device()
	},
}
