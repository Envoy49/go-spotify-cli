package player

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/url"

	"github.com/envoy49/go-spotify-cli/commands"
	"github.com/envoy49/go-spotify-cli/commands/cmdTypes"
	"github.com/envoy49/go-spotify-cli/config"
	"github.com/envoy49/go-spotify-cli/server"
	"github.com/sirupsen/logrus"
	"github.com/spf13/cobra"
)

func playTrack(accessToken, uri, deviceID string) error {
	params := &commands.PlayerParams{
		AccessToken: accessToken,
		Method:      "PUT",
		Endpoint:    spotifyPlayerEndpoint + "/player/play",
	}

	if uri != "" {
		playBody := map[string][]string{"uris": {uri}}
		bodyData, err := json.Marshal(playBody)
		if err != nil {
			return fmt.Errorf("marshal selected track URI: %w", err)
		}
		params.Body = bytes.NewReader(bodyData)
	}

	if deviceID != "" {
		query := url.Values{}
		query.Set("device_id", deviceID)
		params.Endpoint += "?" + query.Encode()
	}

	_, err := commands.Fetch(params)
	return err
}

func Play(cfg *config.Config, accessToken string, URI string) {
	err := playTrack(accessToken, URI, "")

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
	} else if URI == "" {
		Player(cfg)
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
