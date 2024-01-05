package commandTypes

type SavedTracks struct {
	Body
	Items []SavedItem `json:"items"`
}

type SavedItem struct {
	AddedAt string     `json:"added_at"`
	Track   SavedTrack `json:"track"`
}

type SavedTrack struct {
	Album            SavedAlbum        `json:"album"`
	Artists          []SavedArtist     `json:"artists"`
	AvailableMarkets []string          `json:"available_markets"`
	DiscNumber       int               `json:"disc_number"`
	DurationMS       int               `json:"duration_ms"`
	Explicit         bool              `json:"explicit"`
	ExternalIDs      ExternalIDs       `json:"external_ids"`
	ExternalUrls     SavedExternalUrls `json:"external_urls"`
	Href             string            `json:"href"`
	ID               string            `json:"id"`
	IsPlayable       bool              `json:"is_playable"`
	LinkedFrom       interface{}       `json:"linked_from"` // Use interface{} if the type is not known or varies
	Restrictions     SavedRestrictions `json:"restrictions"`
	Name             string            `json:"name"`
	Popularity       int               `json:"popularity"`
	PreviewUrl       string            `json:"preview_url"`
	TrackNumber      int               `json:"track_number"`
	Type             string            `json:"type"`
	Uri              string            `json:"uri"`
	IsLocal          bool              `json:"is_local"`
}

type SavedAlbum struct {
	AlbumType            string            `json:"album_type"`
	TotalTracks          int               `json:"total_tracks"`
	AvailableMarkets     []string          `json:"available_markets"`
	ExternalUrls         SavedExternalUrls `json:"external_urls"`
	Href                 string            `json:"href"`
	ID                   string            `json:"id"`
	Images               []Image           `json:"images"`
	Name                 string            `json:"name"`
	ReleaseDate          string            `json:"release_date"`
	ReleaseDatePrecision string            `json:"release_date_precision"`
	Restrictions         SavedRestrictions `json:"restrictions"`
	Type                 string            `json:"type"`
	Uri                  string            `json:"uri"`
	Artists              []SavedArtist     `json:"artists"`
}

type SavedArtist struct {
	ExternalUrls SavedExternalUrls `json:"external_urls"`
	Followers    SavedFollowers    `json:"followers"`
	Genres       []string          `json:"genres"`
	Href         string            `json:"href"`
	ID           string            `json:"id"`
	Images       []Image           `json:"images"`
	Name         string            `json:"name"`
	Popularity   int               `json:"popularity"`
	Type         string            `json:"type"`
	Uri          string            `json:"uri"`
}

type SavedFollowers struct {
	Href  string `json:"href"`
	Total int    `json:"total"`
}

type SavedRestrictions struct {
	Reason string `json:"reason"`
}

type SavedExternalUrls struct {
	Spotify string `json:"spotify"`
}

type ExternalIDs struct {
	Isrc string `json:"isrc"`
	Ean  string `json:"ean"`
	Upc  string `json:"upc"`
}
