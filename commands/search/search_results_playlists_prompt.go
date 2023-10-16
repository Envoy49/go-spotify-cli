package search

import (
	"fmt"
	"github.com/manifoldco/promptui"
	"github.com/sirupsen/logrus"
	"go-spotify-cli/common"
	"strings"
)

func PlaylistsResultsPrompt(playlists *Playlists) string {
	formattedInfo := make([]string, len(playlists.Items))

	for i, item := range playlists.Items {
		playlistInfo := fmt.Sprintf("%s (Tracks: %d) by %s", item.Name, item.Tracks.Total, item.Owner.DisplayName)
		formattedInfo[i] = playlistInfo
	}

	prompt := promptui.Select{
		Label: "Select Playlist",
		Items: formattedInfo,
		Size:  len(playlists.Items),
		Searcher: func(input string, index int) bool {
			name := formattedInfo[index]
			return strings.Contains(strings.ToLower(name), strings.ToLower(input))
		},
		StartInSearchMode: true,
		Templates: &promptui.SelectTemplates{
			Active:   `{{ "▸" | cyan }} {{ . | yellow | underline }}`,
			Inactive: `{{ " " | faint }} {{ . | faint }}`,
			Selected: `{{ "✔" | green }} {{ . }}`,
			Label:    `{{ ">>" | cyan }} {{ .Label | bold }}`,
			Details:  `{{ "Selected Playlist:" | cyan }} {{ .Name | yellow | bold }}`,
		},
	}

	index, _, err := prompt.Run()
	if err != nil {
		logrus.WithError(err).Error("Prompt failed")
		return ""
	}

	selectedPlaylist := playlists.Items[index]
	formattedPlaylistInfo := fmt.Sprintf(
		"Playlist Name          : %s\n"+
			"Playlist ID            : %s\n"+
			"Total Tracks           : %d\n"+
			"Owner                  : %s\n"+
			"Playlist URI           : %s\n",
		common.ValueStyle.Render(selectedPlaylist.Name),
		common.ValueStyle.Render(selectedPlaylist.ID),
		selectedPlaylist.Tracks.Total,
		common.ValueStyle.Render(selectedPlaylist.Owner.DisplayName),
		common.ValueStyle.Render(selectedPlaylist.URI),
	)

	fullBox := common.BoxStyle.Render(common.HeaderStyle.Render("       SELECTED PLAYLIST INFO         ") + "\n" + formattedPlaylistInfo + "\n")

	fmt.Println(fullBox)

	return selectedPlaylist.URI
}
