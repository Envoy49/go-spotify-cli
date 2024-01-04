package searchPrompt

import (
	"fmt"
	"github.com/manifoldco/promptui"
	"github.com/sirupsen/logrus"
	"github.com/envoy49/go-spotify-cli/common"
	"github.com/envoy49/go-spotify-cli/types"
	"strconv"
	"strings"
)

func AlbumsResultsPrompt(albums *types.Albums) string {
	formattedInfo := make([]string, len(albums.Items))

	for i, item := range albums.Items {
		albumInfo := fmt.Sprintf("%s - %s", item.Artists[0].Name, item.Name)
		formattedInfo[i] = albumInfo
	}

	prompt := promptui.Select{
		Label: "Select an album",
		Items: formattedInfo,
		Size:  len(albums.Items),
		Searcher: func(input string, index int) bool {
			name := formattedInfo[index]
			return strings.Contains(strings.ToLower(name), strings.ToLower(input))
		},
		StartInSearchMode: true, // start the prompt in search mode
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

	selectedAlbum := albums.Items[index]
	formattedAlbumInfo := fmt.Sprintf(
		"Album Name          : %s\n"+
			"Album Type          : %s\n"+
			"Release Date        : %s\n"+
			"Total Tracks        : %s\n"+
			"Album Id            : %s\n"+
			"Album URI           : %s\n",
		common.ValueStyle.Render(selectedAlbum.Name),
		common.ValueStyle.Render(selectedAlbum.AlbumType),
		common.ValueStyle.Render(selectedAlbum.ReleaseDate),
		common.ValueStyle.Render(strconv.Itoa(selectedAlbum.TotalTracks)),
		common.ValueStyle.Render(selectedAlbum.ID),
		common.ValueStyle.Render(selectedAlbum.URI),
	)

	fullBox := common.BoxStyle.Render(common.HeaderStyle.Render("         SELECTED ALBUM INFO          ") + "\n" + formattedAlbumInfo + "\n")

	fmt.Println(fullBox)

	return selectedAlbum.URI
}
