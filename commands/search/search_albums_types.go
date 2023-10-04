package search

import (
	"go-spotify-cli/common"
)

type Albums struct {
	Body
	Items []AlbumItems `json:"items"`
}

type AlbumItems struct {
	AlbumType            string            `json:"album_type"`
	TotalTracks          int               `json:"total_tracks"`
	AvailableMarkets     []string          `json:"available_markets"`
	ExternalURLs         map[string]string `json:"external_urls"`
	Href                 string            `json:"href"`
	ID                   string            `json:"id"`
	Images               []common.Image    `json:"images"`
	Name                 string            `json:"name"`
	ReleaseDate          string            `json:"release_date"`
	ReleaseDatePrecision string            `json:"release_date_precision"`
	Restrictions         AlbumRestrictions `json:"restrictions"`
	Type                 string            `json:"type"`
	URI                  string            `json:"uri"`
	Artists              []AlbumArtist     `json:"artists"`
}

type AlbumRestrictions struct {
	Reason string `json:"reason"`
}

type AlbumExternalURLs struct {
	Spotify string `json:"spotify"`
}

type AlbumArtist struct {
	ExternalURLs AlbumExternalURLs `json:"external_urls"`
	Href         string            `json:"href"`
	ID           string            `json:"id"`
	Name         string            `json:"name"`
	Type         string            `json:"type"`
	URI          string            `json:"uri"`
}
