package prompt

import (
	"encoding/json"
	"go-spotify-cli/types"
	"log"
)

func SpotifySearchResultsPrompt(body []byte) (string, string) {
	var response *types.SpotifySearchResponse
	err := json.Unmarshal(body, &response)
	if err != nil {
		log.Fatalf("Error decoding JSON: %v", err)
	}

	switch {
	case response.Tracks != nil:
		return TracksResultsPrompt(response.Tracks)
	//case response.Episodes != nil:
	//	return EpisodesResultsPrompt(response.Episodes)
	//case response.Albums != nil:
	//	return AlbumsResultsPrompt(response.Albums)
	//case response.Artists != nil:
	//	return ArtistsResultsPrompt(response.Artists)
	//case response.Audiobooks != nil:
	//	return AudiobooksResultsPrompt(response.Audiobooks)
	//case response.Shows != nil:
	//	return ShowsResultsPrompt(response.Shows)
	//case response.Playlists != nil:
	//	return PlaylistsResultsPrompt(response.Playlists)
	default:
		return "", ""
	}

}
