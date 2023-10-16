package search

import (
	"github.com/sirupsen/logrus"
	"github.com/spf13/cobra"
	"go-spotify-cli/commands"
	"go-spotify-cli/commands/player"
	"go-spotify-cli/common"
	"go-spotify-cli/constants"
	"go-spotify-cli/server"
	"net/url"
)

func buildSpotifySearchURL(baseEndpoint string, prompt *SpotifySearchQuery) string {
	values := url.Values{}
	values.Add("q", prompt.Query)
	values.Add("type", prompt.Type)
	values.Add("limit", prompt.Limit)

	fullURL := baseEndpoint + "?" + values.Encode()

	return fullURL
}

func search(accessToken string, prompt *SpotifySearchQuery) {
	params := &commands.PlayerParams{
		AccessToken: accessToken,
		Method:      "GET",
		Endpoint:    buildSpotifySearchURL(constants.SpotifySearchEndpoint, prompt),
	}

	body, err := commands.FetchCommand(params)

	if err != nil {
		switch e := err.(type) {
		case common.SpotifyAPIError:
			if e.Detail.Error.Message == "Player command failed: No active device found" {
				player.Device()
			}
		}

		logrus.WithError(err).Error("Error searching tracks")
	} else {
		URI := SpotifySearchResultsPrompt(body)
		player.Play(accessToken, URI)
		// call Play function after Search Results Prompt
	}
}

var SendSearchCommand = &cobra.Command{
	Use:   "search",
	Short: "Search spotify song",
	Run: func(cmd *cobra.Command, args []string) {
		token := server.ReadUserModifyTokenOrFetchFromServer()
		err, query := SpotifySearchQueryPrompt()
		if err != nil {
			logrus.WithError(err).Error("Error getting Search Query Prompts")
		}

		search(token, query)
	},
}
