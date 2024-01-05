package searchPrompt

import (
	"fmt"
	"github.com/envoy49/go-spotify-cli/commands"
	"github.com/envoy49/go-spotify-cli/types"
	"github.com/manifoldco/promptui"
	"github.com/sirupsen/logrus"
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
		commands.ValueStyle.Render(selectedShow.Name),
		commands.ValueStyle.Render(selectedShow.ID),
		selectedShow.TotalEpisodes,
		commands.ValueStyle.Render(selectedShow.URI),
	)

	fullBox := commands.BoxStyle.Render(commands.HeaderStyle.Render("       SELECTED SHOW INFO         ") + "\n" + formattedShowInfo + "\n")

	fmt.Println(fullBox)

	return selectedShow.URI
}
