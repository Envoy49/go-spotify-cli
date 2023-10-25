package search

import (
	"fmt"
	"github.com/manifoldco/promptui"
	"github.com/sirupsen/logrus"
	"go-spotify-cli/common"
	"strconv"
	"strings"
)

func TracksResultsPrompt(tracks *Tracks) (string, string) {

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

	prompt := promptui.Select{
		Label: "Select track",
		Items: formattedInfo,
		Size:  len(formattedInfo),
		Searcher: func(input string, index int) bool {
			name := formattedInfo[index]
			return strings.Contains(strings.ToLower(name), strings.ToLower(input))
		},
		StartInSearchMode: true,
		Templates: &promptui.SelectTemplates{
			Active:   `{{ "▸" | bold | blue }} {{ . | underline | blue }}`,
			Inactive: `{{if eq . ">>> NEXT >>>"}}{{ " " | faint }} {{ . | green }}{{else if eq . "<<< PREVIOUS <<<"}}{{ " " | faint }} {{ . | red }}{{else}}{{ " " | faint }} {{ . | faint }}{{end}}`,
			Selected: `{{ "✔" | green }} {{ . | cyan }}`,
			Label:    `{{ ">>" | bold | cyan }} {{ .Label | bold }}`,
		},
	}

	index, _, err := prompt.Run()
	if err != nil {
		logrus.WithError(err).Error("Prompt failed")
		return "", ""
	}

	lastIndex := len(tracks.Items)

	if lastIndex == index {
		return "", tracks.Next
	}

	if lastIndex+1 == index {
		return "", tracks.Previous
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

	return selectedTrack.URI, ""
}
