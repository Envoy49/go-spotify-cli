package player

import (
	"github.com/envoy49/go-spotify-cli/commands"
	"github.com/envoy49/go-spotify-cli/constants"
	"github.com/envoy49/go-spotify-cli/types"
	"github.com/sirupsen/logrus"
	"time"
)

func AddToQueue(accessToken string, url string) {
	time.Sleep(1 * time.Second) // add 1-second delay so spotify has time to update previous track
	params := &commands.PlayerParams{
		AccessToken: accessToken,
		Method:      "POST",
		Endpoint:    constants.SpotifyPlayerEndpoint + "/player/queue" + "?uri=" + url,
	}
	_, err := commands.Fetch(params)

	if err != nil {
		switch e := err.(type) {
		case types.SpotifyAPIError:
			if e.Detail.Error.Message == "Player command failed: No active device found" {
				Device()
			}
		default:
			logrus.WithError(err).Error("Error adding item to the queue")
			return
		}
	}
}
