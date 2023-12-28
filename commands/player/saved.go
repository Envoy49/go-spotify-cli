package player

import (
	"encoding/json"
	"fmt"
	"github.com/sirupsen/logrus"
	"github.com/spf13/cobra"
	"go-spotify-cli/commands"
	"go-spotify-cli/common"
	"go-spotify-cli/constants"
	"go-spotify-cli/loader"
	"go-spotify-cli/prompt"
	"go-spotify-cli/server"
	"go-spotify-cli/types"
	"log"
	"strconv"
)

func saved(accessToken string, nextUrl string) *types.SearchPromptResults {
	loader.Start()
	var endpoint = constants.SpotifyPlayerEndpoint + "/tracks"
	if len(nextUrl) > 0 {
		endpoint = nextUrl
	}

	params := &commands.PlayerParams{
		AccessToken: accessToken,
		Method:      "GET",
		Endpoint:    endpoint,
	}
	body, fetchErr := commands.Fetch(params)
	loader.Stop()
	if fetchErr != nil {
		logrus.WithError(fetchErr).Error("Error fetching saved tracks")
	}
	var response *types.SavedTracks

	unmarshalErr := json.Unmarshal(body, &response)

	if unmarshalErr != nil {
		log.Fatalf("Error decoding JSON: %v", unmarshalErr)
		return &types.SearchPromptResults{}
	}

	formattedInfo := make([]string, len(response.Items))

	for i, item := range response.Items {
		trackInfo := fmt.Sprintf("%s - %s", item.Track.Artists[0].Name, item.Track.Name)
		formattedInfo[i] = trackInfo
	}

	if len(response.Next) > 0 {
		formattedInfo = append(formattedInfo, ">>> NEXT >>>")
	}

	if len(response.Previous) > 0 {
		formattedInfo = append(formattedInfo, "<<< PREVIOUS <<<")
	}

	config := &types.SelectionPromptConfig{
		Label:         "Select saved track",
		FormattedInfo: formattedInfo,
	}

	savedPrompt := prompt.CreateSelectionPrompt(config)

	index, _, err := savedPrompt.Run()
	if err != nil {
		logrus.WithError(err).Error("Prompt failed")
		return &types.SearchPromptResults{}
	}

	lastIndex := len(response.Items)

	if lastIndex == index {
		return saved(accessToken, response.Next)
	}

	if lastIndex+1 == index {
		return saved(accessToken, response.Previous)
	}

	selectedTrack := response.Items[index]
	formattedSongInfo := fmt.Sprintf(
		"Artist Name             : %s\n"+
			"Track Name              : %s\n"+
			"Track ID                : %s\n"+
			"Track Popularity        : %s\n"+
			"Track URI               : %s\n",
		common.ValueStyle.Render(selectedTrack.Track.Artists[0].Name),
		common.ValueStyle.Render(selectedTrack.Track.Name),
		common.ValueStyle.Render(selectedTrack.Track.ID),
		common.ValueStyle.Render(strconv.Itoa(selectedTrack.Track.Popularity)),
		common.ValueStyle.Render(selectedTrack.Track.Uri),
	)

	fullBox := common.BoxStyle.Render(common.HeaderStyle.Render("         SELECTED SAVED TRACKs INFO          ") + "\n" + formattedSongInfo + "\n")

	fmt.Println(fullBox)

	return &types.SearchPromptResults{
		PlayUrl: selectedTrack.Track.Uri,
	}
}

var SavedCommand = &cobra.Command{
	Use:   "saved",
	Short: "Saved spotify tracks",
	PreRun: func(cmd *cobra.Command, args []string) {
		loader.Stop()
	},
	Run: func(cmd *cobra.Command, args []string) {
		token := server.ReadUserLibraryReadTokenOrFetchFromServer()
		result := saved(token, "")
		if len(result.PlayUrl) > 0 {
			token := server.ReadUserModifyTokenOrFetchFromServer()
			// instead of Calling Play function, we are adding song to the queue and using Next function
			// otherwise song playing further nexts is not possible, seems like an API limitation.
			//Play(token, result.PlayUrl)
			AddToQueue(token, result.PlayUrl)
			Next(token, false)
		}
	},
}
