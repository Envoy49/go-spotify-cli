package commandTypes

type ArtistsExternalURLs struct {
	Spotify string `json:"spotify"`
}

type ArtistsFollowers struct {
	Href  string `json:"href"`
	Total int    `json:"total"`
}

type ArtistItems struct {
	ExternalURLs ArtistsExternalURLs `json:"external_urls"`
	Followers    ArtistsFollowers    `json:"followers"`
	Genres       []string            `json:"genres"`
	Href         string              `json:"href"`
	ID           string              `json:"id"`
	Images       []Image             `json:"images"`
	Name         string              `json:"name"`
	Popularity   int                 `json:"popularity"`
	Type         string              `json:"type"`
	URI          string              `json:"uri"`
}

type Artists struct {
	Body
	Items []ArtistItems `json:"items"`
}
