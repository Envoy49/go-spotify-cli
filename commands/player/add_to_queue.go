package player

import (
	"github.com/sirupsen/logrus"
	"go-spotify-cli/commands"
	"go-spotify-cli/constants"
	"go-spotify-cli/types"
)

func AddToQueue(accessToken string, url string) {
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
