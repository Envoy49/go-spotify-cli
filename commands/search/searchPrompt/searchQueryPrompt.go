package searchPrompt

import (
	"fmt"
	"log"
	"strconv"
	"strings"

	"github.com/envoy49/go-spotify-cli/commands/cmdTypes"
	"github.com/manifoldco/promptui"
)

// var searchTypes = []string{"Track", "Artist", "Album", "Playlist", "Show", "Episode", "Audiobook"}.
var searchTypes = []string{"Track", "Episode"}

func SpotifySearchQueryPrompt() (error, *cmdTypes.SpotifySearchQuery) {
	promptSearchQuery := promptui.Prompt{
		Label: "Enter your search query",
	}
	searchQuery, err := promptSearchQuery.Run()
	if err != nil {
		log.Fatalf("Prompt failed %v\n", err)
		return err, nil
	}

	promptSearchType := promptui.Select{
		Label: "Select the type(s) of items you want to search for(only Tracks and Episodes are supported now)",
		Items: searchTypes,
		Size:  len(searchTypes),
	}
	_, searchType, err := promptSearchType.Run()
	if err != nil {
		log.Fatalf("Prompt failed %v\n", err)
		return err, nil
	}
	promptLimit := promptui.Prompt{
		Label:   "How many results do you want to see per item type? (Default is 20, max 50)",
		Default: "20",
		Validate: func(input string) error {
			value, err := strconv.Atoi(input)
			if err != nil || value < 0 || value > 50 {
				return fmt.Errorf("enter a valid number between 0 and 50")
			}
			return nil
		},
	}
	limitStr, err := promptLimit.Run()
	if err != nil {
		log.Fatalf("Prompt failed %v\n", err)
		return err, nil
	}

	var result = &cmdTypes.SpotifySearchQuery{
		Query: searchQuery,
		Type:  strings.ToLower(searchType),
		Limit: limitStr,
	}

	return nil, result
}
