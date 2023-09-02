package commands

import (
	"github.com/sirupsen/logrus"
	"github.com/spf13/cobra"
	"go-spotify-cli/cmd/player"
	"go-spotify-cli/server"
)

func device() {
	token := server.FetchDeviceTokenFromBrowser()
	params := &commands.PlayerParams{
		AccessToken: token,
		Method:      "GET",
		Endpoint:    "/devices",
	}
	_, response, _ := commands.Player(params)
	
	logrus.Info("List of available devices:", response)
}

var DeviceCommand = &cobra.Command{
	Use:   "device",
	Short: "Get all connected devices",
	Run: func(cmd *cobra.Command, args []string) {
		device()
	},
}
