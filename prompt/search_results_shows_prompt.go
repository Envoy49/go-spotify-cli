package prompt

import (
	"fmt"
	"github.com/manifoldco/promptui"
	"github.com/sirupsen/logrus"
	"go-spotify-cli/common"
	"go-spotify-cli/types"
	"strings"
)

func ShowsResultsPrompt(shows *types.Shows) string {
	formattedInfo := make([]string, len(shows.Items))

	for i, item := range shows.Items {
		showInfo := fmt.Sprintf("%s (Total Episodes: %d)", item.Name, item.TotalEpisodes)
		formattedInfo[i] = showInfo
	}

	prompt := promptui.Select{
		Label: "Select Show",
		Items: formattedInfo,
		Size:  len(shows.Items),
		Searcher: func(input string, index int) bool {
			name := formattedInfo[index]
			return strings.Contains(strings.ToLower(name), strings.ToLower(input))
		},
		StartInSearchMode: true,
		Templates: &promptui.SelectTemplates{
			Active:   `{{ "▸" | bold | blue }} {{ . | underline | blue }}`,
			Inactive: `{{ " " | faint }} {{ . | faint }}`,
			Selected: `{{ "✔" | green }} {{ . | cyan }}`,
			Label:    `{{ ">>" | bold | cyan }} {{ .Label | bold }}`,
		},
	}

	index, _, err := prompt.Run()
	if err != nil {
		logrus.WithError(err).Error("Prompt failed")
		return ""
	}

	selectedShow := shows.Items[index]
	formattedShowInfo := fmt.Sprintf(
		"Show Name              : %s\n"+
			"Show ID                : %s\n"+
			"Total Episodes         : %d\n"+
			"Show URI               : %s\n",
		common.ValueStyle.Render(selectedShow.Name),
		common.ValueStyle.Render(selectedShow.ID),
		selectedShow.TotalEpisodes,
		common.ValueStyle.Render(selectedShow.URI),
	)

	fullBox := common.BoxStyle.Render(common.HeaderStyle.Render("       SELECTED SHOW INFO         ") + "\n" + formattedShowInfo + "\n")

	fmt.Println(fullBox)

	return selectedShow.URI
}