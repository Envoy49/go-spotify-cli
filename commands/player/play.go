package player

import (
	"bytes"
	"encoding/json"

	"github.com/envoy49/go-spotify-cli/commands"
	"github.com/envoy49/go-spotify-cli/commands/cmdTypes"
	"github.com/envoy49/go-spotify-cli/config"
	"github.com/envoy49/go-spotify-cli/server"
	"github.com/sirupsen/logrus"
	"github.com/spf13/cobra"
)

func Play(cfg *config.Config, accessToken string, URI string) {
	playBody := map[string][]string{
		"uris": {URI},
	}

	bodyData, marshalErr := json.Marshal(playBody)
	if marshalErr != nil {
		logrus.WithError(marshalErr).Error("Error marshaling trackID in player function")
		return
	}

	var params = &commands.PlayerParams{
		AccessToken: accessToken,
		Method:      "PUT",
		Endpoint:    spotifyPlayerEndpoint + "/player/play",
	}

	if len(URI) > 0 {
		params.Body = bytes.NewReader(bodyData)
	}

	_, err := commands.Fetch(params)

	if err != nil {
		switch e := err.(type) {
		case cmdTypes.SpotifyAPIError:
			if e.Detail.Error.Message == "Player command failed: No active device found" {
				Device(cfg)
			}
		default:
			logrus.WithError(err).Error("Some other SpotifyAPIError occurred")
			return
		}
	} else {
		if len(URI) == 0 {
			Player(cfg)
		}
	}
}

func PlayCommand(cfg *config.Config) *cobra.Command {
	return &cobra.Command{
		Use:   "play",
		Short: "Play spotify song",
		Run: func(cmd *cobra.Command, args []string) {
			token := server.ReadUserModifyTokenOrFetchFromServer(cfg)
			Play(cfg, token, "")
		},
	}
}
