package search

import "go-spotify-cli/common"

type Playlists struct {
	Body
	Items []PlaylistItem `json:"items"`
}

type PlaylistItem struct {
	Collaborative bool              `json:"collaborative"`
	Description   string            `json:"description"`
	ExternalURLs  map[string]string `json:"external_urls"`
	Href          string            `json:"href"`
	ID            string            `json:"id"`
	Images        []common.Image    `json:"images"`
	Name          string            `json:"name"`
	Owner         PlaylistOwner     `json:"owner"`
	Public        bool              `json:"public"`
	SnapshotID    string            `json:"snapshot_id"`
	Tracks        PlaylistTrackInfo `json:"tracks"`
	Type          string            `json:"type"`
	URI           string            `json:"uri"`
}

type PlaylistOwner struct {
	ExternalURLs map[string]string      `json:"external_urls"`
	Followers    PlaylistOwnerFollowers `json:"followers"`
	Href         string                 `json:"href"`
	ID           string                 `json:"id"`
	Type         string                 `json:"type"`
	URI          string                 `json:"uri"`
	DisplayName  string                 `json:"display_name"`
}

type PlaylistOwnerFollowers struct {
	Href  string `json:"href"`
	Total int    `json:"total"`
}

type PlaylistTrackInfo struct {
	Href  string `json:"href"`
	Total int    `json:"total"`
}
