package searchPrompt

import (
	"fmt"
	"github.com/sirupsen/logrus"
	"github.com/envoy49/go-spotify-cli/common"
	"github.com/envoy49/go-spotify-cli/prompt"
	"github.com/envoy49/go-spotify-cli/types"
	"strconv"
)

func EpisodesResultsPrompt(episodes *types.Episodes) *types.SearchPromptResults {
	formattedInfo := make([]string, len(episodes.Items))

	for i, item := range episodes.Items {
		episodeInfo := fmt.Sprintf("%s (Duration: %s minutes)", item.Name, strconv.Itoa(item.DurationMS/60000))
		formattedInfo[i] = episodeInfo
	}

	if len(episodes.Next) > 0 {
		formattedInfo = append(formattedInfo, ">>> NEXT >>>")
	}

	if len(episodes.Previous) > 0 {
		formattedInfo = append(formattedInfo, "<<< PREVIOUS <<<")
	}

	config := &types.SelectionPromptConfig{
		Label:         "Select Episode",
		FormattedInfo: formattedInfo,
	}

	selectionPrompt := prompt.CreateSelectionPrompt(config)

	index, _, err := selectionPrompt.Run()
	if err != nil {
		logrus.WithError(err).Error("Prompt failed")
		return &types.SearchPromptResults{}
	}

	lastIndex := len(episodes.Items)

	if lastIndex == index {
		return &types.SearchPromptResults{
			NextUrl: episodes.Next,
		}
	}

	if lastIndex+1 == index {
		return &types.SearchPromptResults{
			NextUrl: episodes.Previous,
		}
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

	return &types.SearchPromptResults{
		PlayUrl: selectedEpisode.URI,
	}
}
