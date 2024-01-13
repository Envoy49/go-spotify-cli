package searchPrompt

import (
	"fmt"
	"github.com/envoy49/go-spotify-cli/commands"
	"github.com/envoy49/go-spotify-cli/commands/cmdTypes"
	"github.com/sirupsen/logrus"
	"golang.org/x/term"
	"os"
	"strconv"
)

func EpisodesResultsPrompt(episodes *cmdTypes.Episodes) *cmdTypes.SearchPromptResults {
	formattedInfo := make([]string, len(episodes.Items))

	// Get the terminal width
	width, _, err := term.GetSize(int(os.Stdout.Fd()))
	if err != nil {
		logrus.WithError(err).Error("Failed to get terminal size")
		width = 80
	}

	const padding = 30
	maxNameLength := width - padding

	for i, item := range episodes.Items {
		episodeName := item.Name
		if len(episodeName) > maxNameLength {
			episodeName = episodeName[:maxNameLength] + "…" // Truncate and add ellipsis
		}
		episodeInfo := fmt.Sprintf("%s (Dur: %s mins)", episodeName, strconv.Itoa(item.DurationMS/60000))
		formattedInfo[i] = episodeInfo
	}

	if len(episodes.Next) > 0 {
		formattedInfo = append(formattedInfo, ">>> NEXT >>>")
	}

	if len(episodes.Previous) > 0 {
		formattedInfo = append(formattedInfo, "<<< PREVIOUS <<<")
	}

	config := &cmdTypes.SelectionPromptConfig{
		Label:         "Select Episode",
		FormattedInfo: formattedInfo,
	}

	selectionPrompt := CreateSelectionPrompt(config)

	index, _, err := selectionPrompt.Run()
	if err != nil {
		logrus.WithError(err).Error("Prompt failed")
		return &cmdTypes.SearchPromptResults{}
	}

	lastIndex := len(episodes.Items)

	if lastIndex == index {
		return &cmdTypes.SearchPromptResults{
			NextUrl: episodes.Next,
		}
	}

	if lastIndex+1 == index {
		return &cmdTypes.SearchPromptResults{
			NextUrl: episodes.Previous,
		}
	}

	selectedEpisode := episodes.Items[index]
	formattedEpisodeInfo := fmt.Sprintf(
		"Episode Name           : %s\n"+
			"Episode ID             : %s\n"+
			"Duration               : %s minutes\n"+
			"Episode URI            : %s\n",
		commands.ValueStyle.Render(selectedEpisode.Name),
		commands.ValueStyle.Render(selectedEpisode.ID),
		commands.ValueStyle.Render(strconv.Itoa(selectedEpisode.DurationMS/60000)),
		commands.ValueStyle.Render(selectedEpisode.URI),
	)

	fullBox := commands.BoxStyle.Render(commands.HeaderStyle.Render("       SELECTED EPISODE INFO        ") + "\n" + formattedEpisodeInfo + "\n")

	fmt.Println(fullBox)

	return &cmdTypes.SearchPromptResults{
		PlayUrl: selectedEpisode.URI,
	}
}
