package search

import (
	"fmt"
	"github.com/manifoldco/promptui"
	"github.com/sirupsen/logrus"
	"go-spotify-cli/common"
	"strconv"
	"strings"
)

func AudiobooksResultsPrompt(audiobooks *Audiobooks) string {
	formattedInfo := make([]string, len(audiobooks.Items))

	for i, item := range audiobooks.Items {
		bookInfo := fmt.Sprintf("%s - %s", item.Authors[0].Name, item.Name)
		formattedInfo[i] = bookInfo
	}

	prompt := promptui.Select{
		Label: "Select Audiobook",
		Items: formattedInfo,
		Size:  len(audiobooks.Items),
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

	selectedAudiobook := audiobooks.Items[index]
	formattedBookInfo := fmt.Sprintf(
		"Author Name             : %s\n"+
			"Audiobook Name          : %s\n"+
			"Audiobook ID            : %s\n"+
			"Audiobook URI           : %s\n"+
			"Total Chapters          : %s\n",
		common.ValueStyle.Render(selectedAudiobook.Authors[0].Name),
		common.ValueStyle.Render(selectedAudiobook.Name),
		common.ValueStyle.Render(selectedAudiobook.ID),
		common.ValueStyle.Render(selectedAudiobook.URI),
		common.ValueStyle.Render(strconv.Itoa(selectedAudiobook.TotalChapters)),
	)

	fullBox := common.BoxStyle.Render(common.HeaderStyle.Render("      SELECTED AUDIOBOOK INFO       ") + "\n" + formattedBookInfo + "\n")

	fmt.Println(fullBox)

	return selectedAudiobook.URI
}
