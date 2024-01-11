package player

import (
	"github.com/envoy49/go-spotify-cli/commands"
	"github.com/envoy49/go-spotify-cli/commands/commandTypes"
	"github.com/envoy49/go-spotify-cli/config"
	"github.com/envoy49/go-spotify-cli/server"
	"github.com/sirupsen/logrus"
	"github.com/spf13/cobra"
)

func previous(cfg *config.Config, accessToken string) {
	params := &commands.PlayerParams{
		AccessToken: accessToken,
		Method:      "POST",
		Endpoint:    spotifyPlayerEndpoint + "/player/previous",
	}
	_, err := commands.Fetch(params)

	if err != nil {
		switch e := err.(type) {
		case commandTypes.SpotifyAPIError:
			if e.Detail.Error.Message == "Player command failed: No active device found" {
				Device(cfg)
			}
		default:
			logrus.WithError(err).Error("Error going to the previous track")
			return
		}
	} else {
		Player(cfg)
	}
}

func PreviousCommand(cfg *config.Config) *cobra.Command {
	return &cobra.Command{
		Use:   "previous",
		Short: "Previous spotify song",
		Run: func(cmd *cobra.Command, args []string) {
			token := server.ReadUserModifyTokenOrFetchFromServer(cfg)
			previous(cfg, token)
		},
	}
}
