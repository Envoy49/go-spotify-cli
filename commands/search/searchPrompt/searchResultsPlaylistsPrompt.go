package searchPrompt

import (
	"fmt"
	"strings"

	"github.com/envoy49/go-spotify-cli/commands"
	"github.com/envoy49/go-spotify-cli/commands/cmdTypes"
	"github.com/manifoldco/promptui"
	"github.com/sirupsen/logrus"
)

func PlaylistsResultsPrompt(playlists *cmdTypes.Playlists) string {
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

	selectedPlaylist := playlists.Items[index]
	formattedPlaylistInfo := fmt.Sprintf(
		"Playlist Name          : %s\n"+
			"Playlist ID            : %s\n"+
			"Total Tracks           : %d\n"+
			"Owner                  : %s\n"+
			"Playlist URI           : %s\n",
		commands.ValueStyle.Render(selectedPlaylist.Name),
		commands.ValueStyle.Render(selectedPlaylist.ID),
		selectedPlaylist.Tracks.Total,
		commands.ValueStyle.Render(selectedPlaylist.Owner.DisplayName),
		commands.ValueStyle.Render(selectedPlaylist.URI),
	)

	fullBox := commands.BoxStyle.Render(commands.HeaderStyle.Render("       SELECTED PLAYLIST INFO         ") + "\n" + formattedPlaylistInfo + "\n")

	fmt.Println(fullBox)

	return selectedPlaylist.URI
}
