package searchPrompt

import (
	"fmt"
	"github.com/envoy49/go-spotify-cli/commands"
	"github.com/envoy49/go-spotify-cli/commands/cmdTypes"
	"github.com/manifoldco/promptui"
	"github.com/sirupsen/logrus"
	"strconv"
	"strings"
)

func AlbumsResultsPrompt(albums *cmdTypes.Albums) string {
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
		commands.ValueStyle.Render(selectedAlbum.Name),
		commands.ValueStyle.Render(selectedAlbum.AlbumType),
		commands.ValueStyle.Render(selectedAlbum.ReleaseDate),
		commands.ValueStyle.Render(strconv.Itoa(selectedAlbum.TotalTracks)),
		commands.ValueStyle.Render(selectedAlbum.ID),
		commands.ValueStyle.Render(selectedAlbum.URI),
	)

	fullBox := commands.BoxStyle.Render(commands.HeaderStyle.Render("         SELECTED ALBUM INFO          ") + "\n" + formattedAlbumInfo + "\n")

	fmt.Println(fullBox)

	return selectedAlbum.URI
}
