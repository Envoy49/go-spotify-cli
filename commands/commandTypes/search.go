package commandTypes

type SpotifySearchResponse struct {
	Tracks     *Tracks     `json:"tracks"`
	Artists    *Artists    `json:"artists"`
	Albums     *Albums     `json:"albums"`
	Playlists  *Playlists  `json:"playlists"`
	Shows      *Shows      `json:"shows"`
	Episodes   *Episodes   `json:"episodes"`
	Audiobooks *Audiobooks `json:"audiobooks"`
}

type SearchPromptResults struct {
	PlayUrl string
	NextUrl string
}

type SpotifySearchQuery struct {
	Query string
	Type  string
	Limit string
}
