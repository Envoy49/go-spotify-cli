package player

import (
	"encoding/json"
	"fmt"
	"github.com/envoy49/go-spotify-cli/commands/commandTypes"
	"github.com/envoy49/go-spotify-cli/commands/search/searchPrompt"
	"github.com/envoy49/go-spotify-cli/config"
	"golang.org/x/term"
	"log"
	"os"
	"strconv"

	"github.com/envoy49/go-spotify-cli/commands"
	"github.com/envoy49/go-spotify-cli/loader"
	"github.com/envoy49/go-spotify-cli/server"
	"github.com/sirupsen/logrus"
	"github.com/spf13/cobra"
)

func saved(accessToken string, nextUrl string) *commandTypes.SearchPromptResults {
	loader.Start()
	var endpoint = spotifyPlayerEndpoint + "/tracks"
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
	var response *commandTypes.SavedTracks

	unmarshalErr := json.Unmarshal(body, &response)

	if unmarshalErr != nil {
		log.Fatalf("Error decoding JSON: %v", unmarshalErr)
		return &commandTypes.SearchPromptResults{}
	}

	formattedInfo := make([]string, len(response.Items))

	// Get the terminal width
	width, _, err := term.GetSize(int(os.Stdout.Fd()))
	if err != nil {
		logrus.WithError(err).Error("Failed to get terminal size")
		width = 80 // Default width in case of error
	}

	const padding = 30
	maxNameLength := width - padding

	for i, item := range response.Items {
		artistName := item.Track.Artists[0].Name
		trackName := item.Track.Name

		combinedName := fmt.Sprintf("%s - %s", artistName, trackName)
		if len(combinedName) > maxNameLength {
			combinedName = combinedName[:maxNameLength] + "…" // Truncate and add ellipsis
		}

		formattedInfo[i] = combinedName
	}

	if len(response.Next) > 0 {
		formattedInfo = append(formattedInfo, ">>> NEXT >>>")
	}

	if len(response.Previous) > 0 {
		formattedInfo = append(formattedInfo, "<<< PREVIOUS <<<")
	}

	config := &commandTypes.SelectionPromptConfig{
		Label:         "Select saved track",
		FormattedInfo: formattedInfo,
	}

	savedPrompt := searchPrompt.CreateSelectionPrompt(config)

	index, _, err := savedPrompt.Run()
	if err != nil {
		logrus.WithError(err).Error("Prompt failed")
		return &commandTypes.SearchPromptResults{}
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
		commands.ValueStyle.Render(selectedTrack.Track.Artists[0].Name),
		commands.ValueStyle.Render(selectedTrack.Track.Name),
		commands.ValueStyle.Render(selectedTrack.Track.ID),
		commands.ValueStyle.Render(strconv.Itoa(selectedTrack.Track.Popularity)),
		commands.ValueStyle.Render(selectedTrack.Track.Uri),
	)

	fullBox := commands.BoxStyle.Render(commands.HeaderStyle.Render("         SELECTED SAVED TRACKs INFO          ") + "\n" + formattedSongInfo + "\n")

	fmt.Println(fullBox)

	return &commandTypes.SearchPromptResults{
		PlayUrl: selectedTrack.Track.Uri,
	}
}

func SavedCommand(cfg *config.Config) *cobra.Command {
	return &cobra.Command{
		Use:   "saved",
		Short: "Saved spotify tracks",
		PreRun: func(cmd *cobra.Command, args []string) {
			loader.Stop()
		},
		Run: func(cmd *cobra.Command, args []string) {
			token := server.ReadUserLibraryReadTokenOrFetchFromServer(cfg)
			result := saved(token, "")
			if len(result.PlayUrl) > 0 {
				token := server.ReadUserModifyTokenOrFetchFromServer(cfg)
				// instead of Calling Play function, we are adding song to the queue and using Next function
				// otherwise song playing further nexts is not possible, seems like an API limitation.
				//Play(token, result.PlayUrl)
				AddToQueue(cfg, token, result.PlayUrl)
				Next(cfg, token, false)
			}
		},
	}
}
