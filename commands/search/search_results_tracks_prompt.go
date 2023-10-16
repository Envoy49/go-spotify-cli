package search

import (
	"fmt"
	"github.com/manifoldco/promptui"
	"github.com/sirupsen/logrus"
	"go-spotify-cli/common"
	"strconv"
	"strings"
)

func TracksResultsPrompt(tracks *Tracks) string {
	formattedInfo := make([]string, len(tracks.Items))

	for i, item := range tracks.Items {
		trackInfo := fmt.Sprintf("%s - %s", item.Artists[0].Name, item.Name)
		formattedInfo[i] = trackInfo
	}

	prompt := promptui.Select{
		Label: "Select track",
		Items: formattedInfo,
		Size:  len(tracks.Items),
		Searcher: func(input string, index int) bool {
			name := formattedInfo[index]
			return strings.Contains(strings.ToLower(name), strings.ToLower(input))
		},
		StartInSearchMode: true, // start the prompt in search mode
		Templates: &promptui.SelectTemplates{
			Active:   `{{ "▸" | cyan }} {{ . | yellow | underline }}`, // underline the active item for emphasis
			Inactive: `{{ " " | faint }} {{ . | faint }}`,
			Selected: `{{ "✔" | green }} {{ . }}`,
			// Also, you can customize how each item appears, separating artist and track name for example
			Label: `{{ ">>" | cyan }} {{ .Label | bold }}`,
			Details: `{{ "Selected Track:" | cyan }}
				 {{ "Artist:" | yellow | bold }} {{ .ArtistName }}
				 {{ "Track:" | yellow | bold }}  {{ .TrackName }}`,
		},
	}

	index, _, err := prompt.Run()
	if err != nil {
		logrus.WithError(err).Error("Prompt failed")
		return ""
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

	return selectedTrack.URI
}
