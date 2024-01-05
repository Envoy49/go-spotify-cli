package searchPrompt

import (
	"fmt"
	"github.com/envoy49/go-spotify-cli/commands"
	"github.com/envoy49/go-spotify-cli/prompt"
	"github.com/envoy49/go-spotify-cli/types"
	"github.com/sirupsen/logrus"
	"golang.org/x/term"
	"os"
	"strconv"
)

func TracksResultsPrompt(tracks *types.Tracks) *types.SearchPromptResults {
	formattedInfo := make([]string, len(tracks.Items))

	// Get the terminal width
	width, _, err := term.GetSize(int(os.Stdout.Fd()))
	if err != nil {
		logrus.WithError(err).Error("Failed to get terminal size")
		width = 80 // Default width in case of error
	}

	const padding = 30 // Adjust this based on your layout needs
	maxNameLength := width - padding

	for i, item := range tracks.Items {
		artistName := item.Artists[0].Name
		trackName := item.Name

		combinedName := fmt.Sprintf("%s - %s", artistName, trackName)
		if len(combinedName) > maxNameLength {
			combinedName = combinedName[:maxNameLength] + "…" // Truncate and add ellipsis
		}

		formattedInfo[i] = combinedName
	}

	if len(tracks.Next) > 0 {
		formattedInfo = append(formattedInfo, ">>> NEXT >>>")
	}

	if len(tracks.Previous) > 0 {
		formattedInfo = append(formattedInfo, "<<< PREVIOUS <<<")
	}

	config := &types.SelectionPromptConfig{
		Label:         "Select track",
		FormattedInfo: formattedInfo,
	}

	selectionPrompt := prompt.CreateSelectionPrompt(config)

	index, _, err := selectionPrompt.Run()
	if err != nil {
		logrus.WithError(err).Error("Prompt failed")
		return &types.SearchPromptResults{}
	}

	lastIndex := len(tracks.Items)

	if lastIndex == index {
		return &types.SearchPromptResults{
			NextUrl: tracks.Next,
		}
	}

	if lastIndex+1 == index {
		return &types.SearchPromptResults{
			NextUrl: tracks.Previous,
		}
	}

	selectedTrack := tracks.Items[index]
	formattedSongInfo := fmt.Sprintf(
		"Artist Name             : %s\n"+
			"Track Name              : %s\n"+
			"Track ID                : %s\n"+
			"Track Popularity        : %s\n"+
			"Track URI               : %s\n",
		commands.ValueStyle.Render(selectedTrack.Artists[0].Name),
		commands.ValueStyle.Render(selectedTrack.Name),
		commands.ValueStyle.Render(selectedTrack.ID),
		commands.ValueStyle.Render(strconv.Itoa(selectedTrack.Popularity)),
		commands.ValueStyle.Render(selectedTrack.URI),
	)

	fullBox := commands.BoxStyle.Render(commands.HeaderStyle.Render("         SELECTED TRACK INFO          ") + "\n" + formattedSongInfo + "\n")

	fmt.Println(fullBox)

	return &types.SearchPromptResults{
		PlayUrl: selectedTrack.URI,
	}
}
