package player

import (
	"errors"
	"net/url"
	"time"

	"github.com/envoy49/go-spotify-cli/commands"
	"github.com/envoy49/go-spotify-cli/commands/cmdTypes"
	"github.com/envoy49/go-spotify-cli/config"
)

const noActiveDeviceMessage = "Player command failed: No active device found"

// AddToQueue returns true when it started the selected track directly after
// activating a device, so callers do not advance an unrelated existing queue.
func AddToQueue(cfg *config.Config, accessToken string, uri string) (bool, error) {
	time.Sleep(1 * time.Second) // give Spotify time to update the previous track
	query := url.Values{}
	query.Set("uri", uri)
	params := &commands.PlayerParams{
		AccessToken: accessToken,
		Method:      "POST",
		Endpoint:    spotifyPlayerEndpoint + "/player/queue?" + query.Encode(),
	}
	_, err := commands.Fetch(params)
	if err == nil {
		return false, nil
	}

	return recoverQueuePlayback(
		err,
		func() (string, error) { return selectDevice(cfg) },
		func(deviceID string, startPlayback bool) error {
			return transferPlayback(accessToken, deviceID, startPlayback)
		},
		func(deviceID string) error { return playTrack(accessToken, uri, deviceID) },
	)
}

func recoverQueuePlayback(
	queueErr error,
	chooseDevice func() (string, error),
	transfer func(deviceID string, startPlayback bool) error,
	playSelectedTrack func(deviceID string) error,
) (bool, error) {
	var spotifyErr cmdTypes.SpotifyAPIError
	if !errors.As(queueErr, &spotifyErr) || spotifyErr.Detail.Error.Message != noActiveDeviceMessage {
		return false, queueErr
	}

	deviceID, err := chooseDevice()
	if err != nil {
		return false, err
	}
	if err = transfer(deviceID, false); err != nil {
		return false, err
	}
	if err = playSelectedTrack(deviceID); err != nil {
		return false, err
	}

	return true, nil
}
