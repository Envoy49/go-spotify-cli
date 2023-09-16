package commands

import (
	"encoding/json"
	"fmt"
	"github.com/sirupsen/logrus"
	"go-spotify-cli/cmd/player"
	"go-spotify-cli/common"
	"go-spotify-cli/server"
	"log"
)

func Player() {
	token := server.ReadUserReadTokenOrFetchFromServer()
	params := &commands.PlayerParams{
		AccessToken: token,
		Method:      "GET",
	}
	var response common.SpotifyResponse
	body, err := commands.FetchCommand(params)

	if err != nil {
		logrus.WithError(err).Error("Error getting current track")
	} else {
		err := json.Unmarshal(body, &response)
		if err != nil {
			log.Fatalf("Error decoding JSON: %v", err)
		}
		// Print out the information
		formattedSongInfo := fmt.Sprintf(
			"\n------------------------ SONG INFORMATION ------------------------\n"+
				"Song                 : %s - %s\n"+
				"Album                : %s\n"+
				"Album Type           : %s\n"+
				"Album Release Date   : %s\n"+
				"-------------------------------------------------------------------\n",
			response.Item.Artists[0].Name,
			response.Item.Name,
			response.Item.Album.Name,
			response.Item.Album.AlbumType,
			response.Item.Album.ReleaseDate,
		)

		logrus.Info(formattedSongInfo)
	}
}
