package search

import (
	"fmt"
	"github.com/manifoldco/promptui"
	"github.com/sirupsen/logrus"
	"go-spotify-cli/common"
	"strconv"
	"strings"
)

func AlbumsResultsPrompt(albums *Albums) string {
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
			Active:   `{{ "▸" | cyan }} {{ . | yellow | underline }}`, // underline the active item for emphasis
			Inactive: `{{ " " | faint }} {{ . | faint }}`,
			Selected: `{{ "✔" | green }} {{ . }}`,
			Label:    `{{ ">>" | cyan }} {{ .Label | bold }}`,
			Details: `{{ "Selected Album:" | cyan }}
				 {{ "Album:" | yellow | bold }} {{ .ArtistName }}
				 {{ "Track:" | yellow | bold }}  {{ .TrackName }}`,
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
