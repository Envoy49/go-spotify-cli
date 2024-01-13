package player

import (
	"encoding/json"
	"fmt"
	"github.com/envoy49/go-spotify-cli/commands/cmdTypes"
	"github.com/envoy49/go-spotify-cli/config"
	"time"

	"github.com/envoy49/go-spotify-cli/commands"
	"github.com/envoy49/go-spotify-cli/server"
	"github.com/sirupsen/logrus"
)

func Player(cfg *config.Config) {
	time.Sleep(1 * time.Second) // add 1-second delay so spotify has time to update previous track

	token := server.ReadUserReadTokenOrFetchFromServer(cfg)
	params := &commands.PlayerParams{
		AccessToken: token,
		Method:      "GET",
		Endpoint:    spotifyPlayerEndpoint + "/player",
	}
	var response cmdTypes.SpotifyResponse
	body, err := commands.Fetch(params)

	if err != nil {
		logrus.WithError(err).Error("Error getting current track")
		return
	} else {
		err := json.Unmarshal(body, &response)
		if err != nil {
			logrus.WithError(err).Error("Error decoding JSON")
			return
		}
		// Print out the information
		formattedSongInfo := fmt.Sprintf(
			"Song                 : %s%s%s\n"+
				"Album                : %s\n"+
				"Album Type           : %s\n"+
				"Album Release Date   : %s\n",
			commands.ValueStyle.Render(response.Item.Artists[0].Name),
			commands.ValueStyle.Render(" - "),
			commands.ValueStyle.Render(response.Item.Name),
			commands.ValueStyle.Render(response.Item.Album.Name),
			commands.ValueStyle.Render(response.Item.Album.AlbumType),
			commands.ValueStyle.Render(response.Item.Album.ReleaseDate),
		)

		fullBox := commands.BoxStyle.Render(commands.HeaderStyle.Render("         SONG INFORMATION          ") + "\n" + formattedSongInfo + "\n")

		fmt.Println(fullBox)
	}
}
