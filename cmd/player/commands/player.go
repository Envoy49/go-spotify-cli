package commands

import (
	"encoding/json"
	"github.com/sirupsen/logrus"
	"go-spotify-cli/cmd/player"
	"go-spotify-cli/common"
	"go-spotify-cli/server"
	"log"
)

func Player() {
	token := server.FetchDeviceTokenFromBrowser()
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
		logrus.Info("----- SONG INFORMATION -----")

		logrus.Info("Song: ", response.Item.Artists[0].Name+" - "+response.Item.Name)
		logrus.Info("Album: ", response.Item.Album.Name)
		logrus.Info("Album Type: ", response.Item.Album.AlbumType)
		logrus.Info("Album Release Date: ", response.Item.Album.ReleaseDate)
	}
}
