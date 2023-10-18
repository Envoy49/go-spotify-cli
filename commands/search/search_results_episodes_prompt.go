package search

import (
	"fmt"
	"github.com/manifoldco/promptui"
	"github.com/sirupsen/logrus"
	"go-spotify-cli/common"
	"strconv"
	"strings"
)

func EpisodesResultsPrompt(episodes *Episodes) string {
	formattedInfo := make([]string, len(episodes.Items))

	for i, item := range episodes.Items {
		episodeInfo := fmt.Sprintf("%s (Duration: %s minutes)", item.Name, strconv.Itoa(item.DurationMS/60000))
		formattedInfo[i] = episodeInfo
	}

	prompt := promptui.Select{
		Label: "Select Episode",
		Items: formattedInfo,
		Size:  len(episodes.Items),
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

	selectedEpisode := episodes.Items[index]
	formattedEpisodeInfo := fmt.Sprintf(
		"Episode Name           : %s\n"+
			"Episode ID             : %s\n"+
			"Duration               : %s minutes\n"+
			"Episode URI            : %s\n",
		common.ValueStyle.Render(selectedEpisode.Name),
		common.ValueStyle.Render(selectedEpisode.ID),
		common.ValueStyle.Render(strconv.Itoa(selectedEpisode.DurationMS/60000)),
		common.ValueStyle.Render(selectedEpisode.URI),
	)

	fullBox := common.BoxStyle.Render(common.HeaderStyle.Render("       SELECTED EPISODE INFO        ") + "\n" + formattedEpisodeInfo + "\n")

	fmt.Println(fullBox)

	return selectedEpisode.URI
}
