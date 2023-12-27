package search_prompt

import (
	"fmt"
	"github.com/manifoldco/promptui"
	"github.com/sirupsen/logrus"
	"go-spotify-cli/common"
	"go-spotify-cli/types"
	"strconv"
	"strings"
)

func ArtistsResultsPrompt(artists *types.Artists) string {
	formattedInfo := make([]string, len(artists.Items))

	for i, item := range artists.Items {
		artistInfo := item.Name // This time, we only need the artist's name
		formattedInfo[i] = artistInfo
	}

	prompt := promptui.Select{
		Label: "Select artist",
		Items: formattedInfo,
		Size:  len(artists.Items),
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

	selectedArtist := artists.Items[index]
	formattedArtistInfo := fmt.Sprintf(
		"Artist Name             : %s\n"+
			"Artist ID               : %s\n"+
			"Artist Popularity       : %s\n"+
			"Artist URI              : %s\n",
		common.ValueStyle.Render(selectedArtist.Name),
		common.ValueStyle.Render(selectedArtist.ID),
		common.ValueStyle.Render(strconv.Itoa(selectedArtist.Popularity)),
		common.ValueStyle.Render(selectedArtist.URI),
	)

	fullBox := common.BoxStyle.Render(common.HeaderStyle.Render("         SELECTED ARTIST INFO          ") + "\n" + formattedArtistInfo + "\n")

	fmt.Println(fullBox)

	return selectedArtist.URI
}
