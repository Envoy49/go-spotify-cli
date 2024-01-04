package player

import (
	"encoding/json"
	"fmt"
	"github.com/envoy49/go-spotify-cli/commands"
	"github.com/envoy49/go-spotify-cli/common"
	"github.com/envoy49/go-spotify-cli/constants"
	"github.com/envoy49/go-spotify-cli/server"
	"github.com/envoy49/go-spotify-cli/types"
	"github.com/sirupsen/logrus"
	"time"
)

func Player() {
	time.Sleep(1 * time.Second) // add 1-second delay so spotify has time to update previous track

	token := server.ReadUserReadTokenOrFetchFromServer()
	params := &commands.PlayerParams{
		AccessToken: token,
		Method:      "GET",
		Endpoint:    constants.SpotifyPlayerEndpoint + "/player",
	}
	var response types.SpotifyResponse
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
			common.ValueStyle.Render(response.Item.Artists[0].Name),
			common.ValueStyle.Render(" - "),
			common.ValueStyle.Render(response.Item.Name),
			common.ValueStyle.Render(response.Item.Album.Name),
			common.ValueStyle.Render(response.Item.Album.AlbumType),
			common.ValueStyle.Render(response.Item.Album.ReleaseDate),
		)

		fullBox := common.BoxStyle.Render(common.HeaderStyle.Render("         SONG INFORMATION          ") + "\n" + formattedSongInfo + "\n")

		fmt.Println(fullBox)
	}
}
