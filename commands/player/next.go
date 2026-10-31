package player

import (
	"github.com/envoy49/go-spotify-cli/commands"
	"github.com/envoy49/go-spotify-cli/config"
	"github.com/envoy49/go-spotify-cli/server"
	"github.com/sirupsen/logrus"
	"github.com/spf13/cobra"
)

func Next(cfg *config.Config, accessToken string, callPlayer bool) {
	params := &commands.PlayerParams{
		AccessToken: accessToken,
		Method:      "POST",
		Endpoint:    spotifyPlayerEndpoint + "/player/next",
	}
	err := runWithDeviceFallback(func() error {
		_, err := commands.Fetch(params)
		return err
	}, func() {
		Device(cfg)
	})
	if err != nil {
		logrus.WithError(err).Error("Error going to the next track")
		return
	}

	if callPlayer {
		Player(cfg)
	}
}

func NextCommand(cfg *config.Config) *cobra.Command {
	return &cobra.Command{
		Use:   "next",
		Short: "Next spotify song",
		Run: func(cmd *cobra.Command, args []string) {
			token := server.ReadUserModifyTokenOrFetchFromServer(cfg)
			Next(cfg, token, true)
		},
	}
}
