package player

import (
	"github.com/envoy49/go-spotify-cli/commands/cmdTypes"
	"github.com/envoy49/go-spotify-cli/config"
	"time"

	"github.com/envoy49/go-spotify-cli/commands"
	"github.com/sirupsen/logrus"
)

func AddToQueue(cfg *config.Config, accessToken string, url string) {
	time.Sleep(1 * time.Second) // add 1-second delay so spotify has time to update previous track
	params := &commands.PlayerParams{
		AccessToken: accessToken,
		Method:      "POST",
		Endpoint:    spotifyPlayerEndpoint + "/player/queue" + "?uri=" + url,
	}
	_, err := commands.Fetch(params)

	if err != nil {
		switch e := err.(type) {
		case cmdTypes.SpotifyAPIError:
			if e.Detail.Error.Message == "Player command failed: No active device found" {
				Device(cfg)
			}
		default:
			logrus.WithError(err).Error("Error adding item to the queue")
			return
		}
	}
}
