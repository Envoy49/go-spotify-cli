package player

import (
	"encoding/json"
	"fmt"
	"github.com/envoy49/go-spotify-cli/commands/commandTypes"
	"github.com/envoy49/go-spotify-cli/config"
	"log"
	"strings"

	"github.com/charmbracelet/lipgloss"
	"github.com/envoy49/go-spotify-cli/commands"
	"github.com/envoy49/go-spotify-cli/server"
	"github.com/manifoldco/promptui"
	"github.com/sirupsen/logrus"
	"github.com/spf13/cobra"
)

type DeviceResponse struct {
	Devices []commandTypes.DeviceType `json:"devices"`
}

func Device(cfg *config.Config) {
	token := server.ReadUserReadTokenOrFetchFromServer(cfg)
	params := &commands.PlayerParams{
		AccessToken: token,
		Method:      "GET",
		Endpoint:    spotifyPlayerEndpoint + "/player/devices",
	}

	var response DeviceResponse
	body, err := commands.Fetch(params)

	if err != nil {
		logrus.WithError(err).Error("Error getting devices")
		return
	}

	err = json.Unmarshal(body, &response)
	if err != nil {
		log.Fatalf("Error decoding JSON: %v", err)
		return
	}

	for _, device := range response.Devices {
		printDeviceInfo(device)
	}

	deviceNames := make([]string, len(response.Devices))
	for i, device := range response.Devices {
		deviceNames[i] = device.Name
	}

	if len(deviceNames) == 0 {
		fmt.Println("No devices available. Please activate/open at least one device.")
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

	ActivateDevice(cfg, selectedDevice.ID)
}

func printDeviceInfo(device commandTypes.DeviceType) {
	volumeRectStyle := lipgloss.NewStyle().Foreground(lipgloss.Color("#51e2f5"))

	var activeStatusColor string
	var activeStatusSymbol string
	if device.IsActive {
		activeStatusColor = "#00FF00" // Green
		activeStatusSymbol = "✔"
	} else {
		activeStatusColor = "#FF0000" // Red
		activeStatusSymbol = "✖"
	}
	activeStatusStyle := lipgloss.NewStyle().Foreground(lipgloss.Color(activeStatusColor)).Bold(true)

	var privateSessionSymbol string
	if device.IsPrivateSession {
		privateSessionSymbol = "🔒"
	} else {
		privateSessionSymbol = "🔓"
	}
	privateSessionStyle := commands.ValueStyle.Render(privateSessionSymbol)

	var typeSymbol string
	if device.Type == "Smartphone" {
		typeSymbol = "📱"
	} else if device.Type == "Computer" {
		typeSymbol = "💻"
	} else {
		typeSymbol = ""
	}
	typeStyle := commands.ValueStyle.Render(typeSymbol)

	// Calculate the number of emojis to represent the volume level
	numEmojis := device.VolumePercent / 10 // Assuming you want 10 emojis to represent 100%

	// Generate the string of emojis representing the volume level
	volumeEmojis := volumeRectStyle.Render(strings.Repeat("▓", numEmojis)) + strings.Repeat("░", 10-numEmojis)

	formattedInfo := fmt.Sprintf(
		"Device Name       : %s\n"+
			"Is Active         : %s %s\n"+
			"ID                : %s\n"+
			"Private Session   : %s %v\n"+
			"Is Restricted     : %s\n"+
			"Supports Volume   : %s\n"+
			"Type              : %s %s\n"+
			"Volume Percent    : %d%% %s\n",
		commands.ValueStyle.Render(device.Name),
		activeStatusStyle.Render(activeStatusSymbol),
		commands.ValueStyle.Render(fmt.Sprintf("%v", device.IsActive)),
		commands.ValueStyle.Render(device.ID),
		privateSessionStyle,
		commands.ValueStyle.Render(fmt.Sprintf("%v", device.IsPrivateSession)),
		commands.ValueStyle.Render(fmt.Sprintf("%v", device.IsRestricted)),
		commands.ValueStyle.Render(fmt.Sprintf("%v", device.SupportsVolume)),
		typeStyle,
		commands.ValueStyle.Render(device.Type),
		device.VolumePercent,
		volumeEmojis,
	)

	// Combine header and formatted info inside a box
	fullBox := commands.BoxStyle.Render(commands.HeaderStyle.Render("         Device Information          ") + "\n" + formattedInfo + "\n")

	fmt.Println(fullBox)
}

func DeviceCommand(cfg *config.Config) *cobra.Command {
	return &cobra.Command{
		Use:   "device",
		Short: "Get all connected devices",
		Run: func(cmd *cobra.Command, args []string) {
			Device(cfg)
		},
	}
}
