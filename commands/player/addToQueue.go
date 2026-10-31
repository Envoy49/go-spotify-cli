package player

import (
	"time"

	"github.com/envoy49/go-spotify-cli/config"

	"github.com/envoy49/go-spotify-cli/commands"
	"github.com/sirupsen/logrus"
)

func AddToQueue(cfg *config.Config, accessToken string, url string) bool {
	queueTrack := func() error {
		time.Sleep(1 * time.Second) // Give Spotify time to update the current track.
		params := &commands.PlayerParams{
			AccessToken: accessToken,
			Method:      "POST",
			Endpoint:    spotifyPlayerEndpoint + "/player/queue" + "?uri=" + url,
		}
		_, err := commands.Fetch(params)
		return err
	}

	err := runWithDeviceFallback(queueTrack, func() {
		Device(cfg)
	})
	if err != nil {
		logrus.WithError(err).Error("Error adding item to the queue")
		return false
	}

	return true
}
