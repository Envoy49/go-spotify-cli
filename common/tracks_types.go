package common

type SavedTracksResponse struct {
	Tracks []Track `json:"tracks"`
}

type Track struct {
	Album            TracksAlbum    `json:"album"`
	Artists          []TracksArtist `json:"artists"`
	AvailableMarkets []string       `json:"available_markets"`
	DiscNumber       int            `json:"disc_number"`
	DurationMs       int            `json:"duration_ms"`
	Explicit         bool           `json:"explicit"`
	ExternalIDs      ExternalID     `json:"external_ids"`
	ExternalURLs     ExternalURL    `json:"external_urls"`
	Href             string         `json:"href"`
	ID               string         `json:"id"`
	IsPlayable       bool           `json:"is_playable"`
	LinkedFrom       struct{}       `json:"linked_from"`
	Restrictions     Restriction    `json:"restrictions"`
	Name             string         `json:"name"`
	Popularity       int            `json:"popularity"`
	PreviewURL       string         `json:"preview_url"`
	TrackNumber      int            `json:"track_number"`
	Type             string         `json:"type"`
	URI              string         `json:"uri"`
	IsLocal          bool           `json:"is_local"`
}

type TracksAlbum struct {
	AlbumType            string         `json:"album_type"`
	TotalTracks          int            `json:"total_tracks"`
	AvailableMarkets     []string       `json:"available_markets"`
	ExternalURLs         ExternalURL    `json:"external_urls"`
	Href                 string         `json:"href"`
	ID                   string         `json:"id"`
	Images               []Image        `json:"images"`
	Name                 string         `json:"name"`
	ReleaseDate          string         `json:"release_date"`
	ReleaseDatePrecision string         `json:"release_date_precision"`
	Restrictions         Restriction    `json:"restrictions"`
	Type                 string         `json:"type"`
	URI                  string         `json:"uri"`
	Artists              []TracksArtist `json:"artists"`
}

type TracksArtist struct {
	ExternalURLs ExternalURL `json:"external_urls"`
	Followers    Follower    `json:"followers,omitempty"`
	Genres       []string    `json:"genres,omitempty"`
	Href         string      `json:"href"`
	ID           string      `json:"id"`
	Images       []Image     `json:"images,omitempty"`
	Name         string      `json:"name"`
	Popularity   int         `json:"popularity,omitempty"`
	Type         string      `json:"type"`
	URI          string      `json:"uri"`
}

type Follower struct {
	Href  string `json:"href"`
	Total int    `json:"total"`
}

type ExternalURL struct {
	Spotify string `json:"spotify"`
}

type ExternalID struct {
	Isrc string `json:"isrc"`
	EAN  string `json:"ean"`
	UPC  string `json:"upc"`
}

type Restriction struct {
	Reason string `json:"reason"`
}
