package commandTypes

type SpotifyResponse struct {
	Device               DeviceType `json:"device"`
	ShuffleState         bool       `json:"shuffle_state"`
	RepeatState          string     `json:"repeat_state"`
	Timestamp            int64      `json:"timestamp"`
	Context              Context    `json:"context"`
	ProgressMS           int        `json:"progress_ms"`
	Item                 Item       `json:"item"`
	CurrentlyPlayingType string     `json:"currently_playing_type"`
	Actions              Actions    `json:"actions"`
	IsPlaying            bool       `json:"is_playing"`
}

type DeviceType struct {
	Name             string `json:"name"`
	IsActive         bool   `json:"is_active"`
	ID               string `json:"id"`
	IsPrivateSession bool   `json:"is_private_session"`
	IsRestricted     bool   `json:"is_restricted"`
	SupportsVolume   bool   `json:"supports_volume"`
	Type             string `json:"type"`
	VolumePercent    int    `json:"volume_percent"`
}

type Context struct {
	ExternalURLs map[string]string `json:"external_urls"`
	Href         string            `json:"href"`
	Type         string            `json:"type"`
	URI          string            `json:"uri"`
}

type Item struct {
	Album            PlayerAlbum       `json:"album"`
	Artists          []PlayerArtist    `json:"artists"`
	AvailableMarkets []string          `json:"available_markets"`
	DiscNumber       int               `json:"disc_number"`
	DurationMS       int               `json:"duration_ms"`
	Explicit         bool              `json:"explicit"`
	ExternalIDs      map[string]string `json:"external_ids"`
	ExternalURLs     map[string]string `json:"external_urls"`
	Href             string            `json:"href"`
	ID               string            `json:"id"`
	IsLocal          bool              `json:"is_local"`
	Name             string            `json:"name"`
	Popularity       int               `json:"popularity"`
	PreviewURL       string            `json:"preview_url"`
	TrackNumber      int               `json:"track_number"`
	Type             string            `json:"type"`
	URI              string            `json:"uri"`
}

type PlayerAlbum struct {
	AlbumType            string            `json:"album_type"`
	Artists              []PlayerArtist    `json:"artists"`
	AvailableMarkets     []string          `json:"available_markets"`
	ExternalURLs         map[string]string `json:"external_urls"`
	Href                 string            `json:"href"`
	ID                   string            `json:"id"`
	Images               []Image           `json:"images"`
	Name                 string            `json:"name"`
	ReleaseDate          string            `json:"release_date"`
	ReleaseDatePrecision string            `json:"release_date_precision"`
	TotalTracks          int               `json:"total_tracks"`
	Type                 string            `json:"type"`
	URI                  string            `json:"uri"`
}

type PlayerArtist struct {
	ExternalURLs map[string]string `json:"external_urls"`
	Href         string            `json:"href"`
	ID           string            `json:"id"`
	Name         string            `json:"name"`
	Type         string            `json:"type"`
	URI          string            `json:"uri"`
}

type Actions struct {
	Disallows Disallows `json:"disallows"`
}

type Disallows struct {
	Pausing bool `json:"pausing"`
}
