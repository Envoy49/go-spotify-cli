package commands

import (
	"encoding/json"
	"fmt"
	"github.com/sirupsen/logrus"
	"github.com/spf13/cobra"
	"go-spotify-cli/cmd/player"
	"go-spotify-cli/common"
	"go-spotify-cli/server"
	"log"
)

type DeviceResponse struct {
	Devices []common.Device `json:"devices"`
}

func device() {
	token := server.ReadUserReadTokenOrFetchFromServer()
	params := &commands.PlayerParams{
		AccessToken: token,
		Method:      "GET",
		Endpoint:    "/devices",
	}

	var response DeviceResponse
	body, err := commands.FetchCommand(params)

	if err != nil {
		logrus.WithError(err).Error("Error getting devices")
	} else {
		err := json.Unmarshal(body, &response)
		if err != nil {
			log.Fatalf("Error decoding JSON: %v", err)
		}
		// Print out the information

		for ind, device := range response.Devices {
			formattedInfo := fmt.Sprintf(
				"\n------------------------ Device Information: Index %d ------------------------\n"+
					"Device Name       : %s\n"+
					"Is Active         : %v\n"+
					"ID                : %s\n"+
					"Private Session   : %v\n"+
					"Is Restricted     : %v\n"+
					"Supports Volume   : %v\n"+
					"Type              : %s\n"+
					"Volume Percent    : %d\n"+
					"------------------------------------------------------------------------------\n",
				ind,
				device.Name,
				device.IsActive,
				device.ID,
				device.IsPrivateSession,
				device.IsRestricted,
				device.SupportsVolume,
				device.Type,
				device.VolumePercent,
			)
			logrus.Info(formattedInfo)
		}

		//logrus.Info("Song: ", response.Item.Artists[0].Name+" - "+response.Item.Name)
		//logrus.Info("Album: ", response.Item.Album.Name)
		//logrus.Info("Album Type: ", response.Item.Album.AlbumType)
		//logrus.Info("Album Release Date: ", response.Item.Album.ReleaseDate)
	}
}

var DeviceCommand = &cobra.Command{
	Use:   "device",
	Short: "Get all connected devices",
	Run: func(cmd *cobra.Command, args []string) {
		device()
	},
}
