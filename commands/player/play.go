package player

import (
	"bytes"
	"encoding/json"
	"github.com/sirupsen/logrus"
	"github.com/spf13/cobra"
	"go-spotify-cli/commands"
	"go-spotify-cli/common"
	"go-spotify-cli/constants"
	"go-spotify-cli/server"
)

func Play(accessToken string, URI string) {
	playBody := map[string][]string{
		"uris": {URI},
	}

	bodyData, marshalErr := json.Marshal(playBody)
	if marshalErr != nil {
		logrus.WithError(marshalErr).Error("Error marshaling trackID in player function")
	}

	var params = &commands.PlayerParams{
		AccessToken: accessToken,
		Method:      "PUT",
		Endpoint:    constants.SpotifyPlayerEndpoint + "/player/play",
	}

	if len(URI) > 0 {
		params.Body = bytes.NewReader(bodyData)
	}

	_, err := commands.FetchCommand(params)

	if err != nil {
		switch e := err.(type) {
		case common.SpotifyAPIError:
			if e.Detail.Error.Message == "Player command failed: No active device found" {
				Device()
				_, err2 := commands.FetchCommand(params)
				if err2 != nil {
					logrus.WithError(err).Error("Error playing your track")
				}
			} else {
				// Some other action for other SpotifyAPIError types.
				logrus.WithError(err).Error("Some other SpotifyAPIError occurred")
			}
		default:
			logrus.WithError(err).Error("Error playing your track")
		}
	} else {
		logrus.Info("Playing")
		Player()
	}
}

var PlayCommand = &cobra.Command{
	Use:   "play",
	Short: "Play spotify song",
	Run: func(cmd *cobra.Command, args []string) {
		token := server.ReadUserModifyTokenOrFetchFromServer()
		Play(token, "")
	},
}
