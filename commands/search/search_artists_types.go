package search

import (
	"go-spotify-cli/common"
)

type ArtistsExternalURLs struct {
	Spotify string `json:"spotify"`
}

type ArtistsFollowers struct {
	Href  string `json:"href"`
	Total int    `json:"total"`
}

type ArtistsItems struct {
	ExternalURLs ArtistsExternalURLs `json:"external_urls"`
	Followers    ArtistsFollowers    `json:"followers"`
	Genres       []string            `json:"genres"`
	Href         string              `json:"href"`
	ID           string              `json:"id"`
	Images       []common.Image      `json:"images"`
	Name         string              `json:"name"`
	Popularity   int                 `json:"popularity"`
	Type         string              `json:"type"`
	URI          string              `json:"uri"`
}

type Artists struct {
	Body
	Items []ArtistsItems `json:"items"`
}
