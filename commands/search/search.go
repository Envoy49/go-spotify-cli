package search

import (
	"net/url"

	"github.com/envoy49/go-spotify-cli/commands"
	"github.com/envoy49/go-spotify-cli/commands/player"
	"github.com/envoy49/go-spotify-cli/commands/search/searchPrompt"
	"github.com/envoy49/go-spotify-cli/loader"
	"github.com/envoy49/go-spotify-cli/server"
	"github.com/envoy49/go-spotify-cli/types"
	"github.com/sirupsen/logrus"
	"github.com/spf13/cobra"
)

const (
	spotifySearchEndpoint = "https://api.spotify.com/v1/search"
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
	loader.Start()
	var endpoint string
	if query != nil {
		endpoint = buildSpotifySearchURL(spotifySearchEndpoint, query)
	} else {
		endpoint = nextUrl
	}

	params := &commands.PlayerParams{
		AccessToken: accessToken,
		Method:      "GET",
		Endpoint:    endpoint,
	}

	body, err := commands.Fetch(params)
	loader.Stop()

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
		result := searchPrompt.SpotifySearchResultsPrompt(body)
		if len(result.NextUrl) > 0 {
			search(accessToken, nil, result.NextUrl)
		}
		if len(result.PlayUrl) > 0 {
			// instead of Calling Play function, we are adding song to the queue and using Next function
			// otherwise song playing further nexts is not possible
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
		loader.Stop()
		token := server.ReadUserModifyTokenOrFetchFromServer()
		err, query := searchPrompt.SpotifySearchQueryPrompt()
		if err != nil {
			logrus.WithError(err).Error("Error getting Search Query Prompts")
			return
		}

		search(token, query, "")
	},
}
