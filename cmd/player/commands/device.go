package commands

import (
	"encoding/json"
	"fmt"
	"github.com/manifoldco/promptui"
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

func Device() {
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
			fmt.Println(formattedInfo)
		}

		deviceNames := make([]string, len(response.Devices))
		for i, device := range response.Devices {
			deviceNames[i] = device.Name
		}

		if len(deviceNames) == 0 {
			fmt.Println("No devices available. Please activate at least one device.")
			return
		}

		prompt := promptui.Select{
			Label: "Select device to play a track",
			Items: deviceNames,
		}

		selectedIndex, _, err := prompt.Run()
		if err != nil {
			logrus.WithError(err).Error("Prompt failed")
			return
		}

		selectedDevice := response.Devices[selectedIndex]

		ActivateDevice(selectedDevice.ID)
	}
}

var DeviceCommand = &cobra.Command{
	Use:   "device",
	Short: "Get all connected devices",
	Run: func(cmd *cobra.Command, args []string) {
		Device()
	},
}
