package search

import (
	"github.com/sirupsen/logrus"
	"github.com/spf13/cobra"
	"go-spotify-cli/commands"
	"go-spotify-cli/commands/player"
	"go-spotify-cli/constants"
	"go-spotify-cli/prompt/search_prompt"
	"go-spotify-cli/server"
	"go-spotify-cli/spinnerInstance"
	"go-spotify-cli/types"
	"net/url"
)

func buildSpotifySearchURL(baseEndpoint string, prompt *types.SpotifySearchQuery) string {
	values := url.Values{}
	values.Add("q", prompt.Query)
	values.Add("type", prompt.Type)
	values.Add("limit", prompt.Limit)

	fullURL := baseEndpoint + "?" + values.Encode()

	return fullURL
}

func search(accessToken string, query *types.SpotifySearchQuery, nextUrl string) {
	var endpoint string
	if query != nil {
		endpoint = buildSpotifySearchURL(constants.SpotifySearchEndpoint, query)
	} else {
		endpoint = nextUrl
	}

	params := &commands.PlayerParams{
		AccessToken: accessToken,
		Method:      "GET",
		Endpoint:    endpoint,
	}

	body, err := commands.Fetch(params)

	if err != nil {
		switch e := err.(type) {
		case types.SpotifyAPIError:
			if e.Detail.Error.Message == "Player command failed: No active device found" {
				player.Device()
			}
		default:
			logrus.WithError(err).Error("Error searching tracks")
			return
		}

	} else {
		result := search_prompt.SpotifySearchResultsPrompt(body)
		if len(result.NextUrl) > 0 {
			search(accessToken, nil, result.NextUrl)
		}
		if len(result.PlayUrl) > 0 {
			// call Play function after Search Results Prompt
			//player.Play(accessToken, playUrl)
			player.AddToQueue(accessToken, result.PlayUrl)
			player.Next(accessToken, false)
		}
	}
}

var SendSearchCommand = &cobra.Command{
	Use:   "search",
	Short: "Search spotify song",
	Run: func(cmd *cobra.Command, args []string) {
		spinnerInstance.Stop()
		token := server.ReadUserModifyTokenOrFetchFromServer()
		err, query := search_prompt.SpotifySearchQueryPrompt()
		if err != nil {
			logrus.WithError(err).Error("Error getting Search Query Prompts")
			return
		}

		search(token, query, "")
	},
}
