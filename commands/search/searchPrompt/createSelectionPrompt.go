package searchPrompt

import (
	"github.com/envoy49/go-spotify-cli/commands/commandTypes"
	"github.com/manifoldco/promptui"
	"strings"
)

func CreateSelectionPrompt(config *commandTypes.SelectionPromptConfig) promptui.Select {
	return promptui.Select{
		Label: config.Label,
		Items: config.FormattedInfo,
		Size:  len(config.FormattedInfo),
		Searcher: func(input string, index int) bool {
			name := config.FormattedInfo[index]
			return strings.Contains(strings.ToLower(name), strings.ToLower(input))
		},
		StartInSearchMode: true,
		Templates: &promptui.SelectTemplates{
			Active:   `{{ "▸" | bold | blue }} {{ . | underline | blue }}`,
			Inactive: `{{if eq . ">>> NEXT >>>"}}{{ " " | faint }} {{ . | green }}{{else if eq . "<<< PREVIOUS <<<"}}{{ " " | faint }} {{ . | red }}{{else}}{{ " " | faint }} {{ . | faint }}{{end}}`,
			Selected: `{{ "✔" | green }} {{ . | cyan }}`,
			Label:    `{{ ">>" | bold | cyan }} {{ .Label | bold }}`,
		},
	}
}
