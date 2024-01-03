package searchPrompt

import (
	"fmt"
	"github.com/sirupsen/logrus"
	"go-spotify-cli/common"
	"go-spotify-cli/prompt"
	"go-spotify-cli/types"
	"strconv"
)

func TracksResultsPrompt(tracks *types.Tracks) *types.SearchPromptResults {
	formattedInfo := make([]string, len(tracks.Items))

	for i, item := range tracks.Items {
		trackInfo := fmt.Sprintf("%s - %s", item.Artists[0].Name, item.Name)
		formattedInfo[i] = trackInfo
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
		common.ValueStyle.Render(selectedTrack.Artists[0].Name),
		common.ValueStyle.Render(selectedTrack.Name),
		common.ValueStyle.Render(selectedTrack.ID),
		common.ValueStyle.Render(strconv.Itoa(selectedTrack.Popularity)),
		common.ValueStyle.Render(selectedTrack.URI),
	)

	fullBox := common.BoxStyle.Render(common.HeaderStyle.Render("         SELECTED TRACK INFO          ") + "\n" + formattedSongInfo + "\n")

	fmt.Println(fullBox)

	return &types.SearchPromptResults{
		PlayUrl: selectedTrack.URI,
	}
}
