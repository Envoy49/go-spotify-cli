package player

import "github.com/envoy49/go-spotify-cli/commands/cmdTypes"

const noActiveDeviceMessage = "Player command failed: No active device found"

func runWithDeviceFallback(command func() error, activateDevice func()) error {
	err := command()
	apiError, isSpotifyAPIError := err.(cmdTypes.SpotifyAPIError)
	if !isSpotifyAPIError || apiError.Detail.Error.Message != noActiveDeviceMessage {
		return err
	}

	activateDevice()
	return command()
}
