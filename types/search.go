package types

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

type Body struct {
	Href     string `json:"href"`
	Limit    int    `json:"limit"`
	Next     string `json:"next"`
	Offset   int    `json:"offset"`
	Previous string `json:"previous"`
	Total    int    `json:"total"`
}

type SpotifySearchQuery struct {
	Query string
	Type  string
	Limit string
}
