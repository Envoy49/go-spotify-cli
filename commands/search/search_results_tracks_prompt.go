package search

import (
	"fmt"
	"github.com/manifoldco/promptui"
	"github.com/sirupsen/logrus"
	"go-spotify-cli/common"
	"strconv"
	"strings"
)

func TracksResultsPrompt(tracks *Tracks, nextUrl string) (string, string) {
	formattedInfo := make([]string, len(tracks.Items))

	for i, item := range tracks.Items {
		trackInfo := fmt.Sprintf("%s - %s", item.Artists[0].Name, item.Name)
		formattedInfo[i] = trackInfo
	}

	formattedInfo = append(formattedInfo, "Get more search results")

	prompt := promptui.Select{
		Label: "Select track",
		Items: formattedInfo,
		Size:  len(tracks.Items) + 1,
		Searcher: func(input string, index int) bool {
			name := formattedInfo[index]
			return strings.Contains(strings.ToLower(name), strings.ToLower(input))
		},
		StartInSearchMode: true,
		Templates: &promptui.SelectTemplates{
			Active:   `{{ "▸" | bold | blue }} {{ . | underline | blue }}`,
			Inactive: `{{if eq . "Get more search results"}}{{ " " | faint }} {{ . | green }}{{else}}{{ " " | faint }} {{ . | faint }}{{end}}`,
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
		return "", nextUrl
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
