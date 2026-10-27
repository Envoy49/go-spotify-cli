package player

import (
	"encoding/json"
	"fmt"
	"runtime"
	"strings"

	"github.com/envoy49/go-spotify-cli/commands/cmdTypes"
	"github.com/envoy49/go-spotify-cli/config"

	"github.com/charmbracelet/lipgloss"
	"github.com/envoy49/go-spotify-cli/commands"
	"github.com/envoy49/go-spotify-cli/server"
	"github.com/manifoldco/promptui"
	"github.com/sirupsen/logrus"
	"github.com/spf13/cobra"
)

type DeviceResponse struct {
	Devices []cmdTypes.DeviceType `json:"devices"`
}

func Device(cfg *config.Config) {
	deviceID, err := selectDevice(cfg)
	if err != nil {
		logrus.WithError(err).Error("Error selecting device")
		return
	}

	ActivateDevice(cfg, deviceID)
}

func selectDevice(cfg *config.Config) (string, error) {
	token := server.ReadUserReadTokenOrFetchFromServer(cfg)
	params := &commands.PlayerParams{
		AccessToken: token,
		Method:      "GET",
		Endpoint:    spotifyPlayerEndpoint + "/player/devices",
	}

	var response DeviceResponse
	body, err := commands.Fetch(params)

	if err != nil {
		return "", fmt.Errorf("getting devices: %w", err)
	}

	err = json.Unmarshal(body, &response)
	if err != nil {
		return "", fmt.Errorf("decoding devices response: %w", err)
	}

	for _, device := range response.Devices {
		printDeviceInfo(device)
	}

	deviceNames := make([]string, len(response.Devices))
	for i, device := range response.Devices {
		deviceNames[i] = device.Name
	}

	if len(deviceNames) == 0 {
		return "", fmt.Errorf("no devices available; activate or open a Spotify device first")
	}

	prompt := promptui.Select{
		Label: "Select device to play a track",
		Items: deviceNames,
	}

	selectedIndex, _, err := prompt.Run()
	if err != nil {
		return "", fmt.Errorf("device selection prompt failed: %w", err)
	}

	selectedDevice := response.Devices[selectedIndex]
	return selectedDevice.ID, nil
}

func printDeviceInfo(device cmdTypes.DeviceType) {
	windowsConsole := runtime.GOOS == "windows"
	volumeRectStyle := lipgloss.NewStyle().Foreground(lipgloss.Color("#51e2f5"))

	var activeStatusColor string
	var activeStatusSymbol string
	if device.IsActive {
		activeStatusColor = "#00FF00" // Green
		activeStatusSymbol = "✔"
		if windowsConsole {
			activeStatusSymbol = "OK"
		}
	} else {
		activeStatusColor = "#FF0000" // Red
		activeStatusSymbol = "✖"
		if windowsConsole {
			activeStatusSymbol = "NO"
		}
	}
	activeStatusStyle := lipgloss.NewStyle().Foreground(lipgloss.Color(activeStatusColor)).Bold(true)

	var privateSessionSymbol string
	if device.IsPrivateSession {
		privateSessionSymbol = "🔒"
		if windowsConsole {
			privateSessionSymbol = "LOCKED"
		}
	} else {
		privateSessionSymbol = "🔓"
		if windowsConsole {
			privateSessionSymbol = "OPEN"
		}
	}
	privateSessionStyle := commands.ValueStyle.Render(privateSessionSymbol)

	var typeSymbol string
	if device.Type == "Smartphone" {
		typeSymbol = "📱"
		if windowsConsole {
			typeSymbol = "PHONE"
		}
	} else if device.Type == "Computer" {
		typeSymbol = "💻"
		if windowsConsole {
			typeSymbol = "PC"
		}
	} else {
		typeSymbol = ""
	}
	typeStyle := commands.ValueStyle.Render(typeSymbol)

	// Calculate the number of filled segments to represent the volume level.
	filledBars := device.VolumePercent / 10

	filledSegment := "▓"
	emptySegment := "░"
	if windowsConsole {
		// ASCII fallbacks avoid missing block glyphs in classic Windows consoles.
		filledSegment = "#"
		emptySegment = "-"
	}
	volumeMeter := volumeRectStyle.Render(strings.Repeat(filledSegment, filledBars)) + strings.Repeat(emptySegment, 10-filledBars)

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
		volumeMeter,
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
