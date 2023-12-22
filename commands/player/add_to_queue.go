package player

import (
	"github.com/sirupsen/logrus"
	"go-spotify-cli/commands"
	"go-spotify-cli/common"
	"go-spotify-cli/constants"
)

func AddToQueue(accessToken string, url string) {
	params := &commands.PlayerParams{
		AccessToken: accessToken,
		Method:      "POST",
		Endpoint:    constants.SpotifyPlayerEndpoint + "/player/queue" + "?uri=" + url,
	}
	_, err := commands.FetchCommand(params)

	if err != nil {
		switch e := err.(type) {
		case common.SpotifyAPIError:
			if e.Detail.Error.Message == "Player command failed: No active device found" {
				Device()
			}
		default:
			logrus.WithError(err).Error("Error adding item to the queue")
			return
		}

	}
}
