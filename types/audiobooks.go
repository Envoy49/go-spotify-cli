package types

type Audiobooks struct {
	Body
	Items []AudiobookItems `json:"items"`
}

type AudiobookItems struct {
	Authors          []AudiobookAuthor    `json:"authors"`
	AvailableMarkets []string             `json:"available_markets"`
	Copyrights       []AudiobookCopyright `json:"copyrights"`
	Description      string               `json:"description"`
	HTMLDescription  string               `json:"html_description"`
	Edition          string               `json:"edition"`
	Explicit         bool                 `json:"explicit"`
	ExternalURLs     map[string]string    `json:"external_urls"`
	Href             string               `json:"href"`
	ID               string               `json:"id"`
	Images           []Image              `json:"images"`
	Languages        []string             `json:"languages"`
	MediaType        string               `json:"media_type"`
	Name             string               `json:"name"`
	Narrators        []AudiobookNarrator  `json:"narrators"`
	Publisher        string               `json:"publisher"`
	Type             string               `json:"type"`
	URI              string               `json:"uri"`
	TotalChapters    int                  `json:"total_chapters"`
}

type AudiobookAuthor struct {
	Name string `json:"name"`
}

type AudiobookNarrator struct {
	Name string `json:"name"`
}

type AudiobookCopyright struct {
	Text string `json:"text"`
	Type string `json:"type"`
}
