package types

type Shows struct {
	Body
	Items []ShowItem `json:"items"`
}

type ShowItem struct {
	AvailableMarkets   []string          `json:"available_markets"`
	Copyrights         []ShowCopyright   `json:"copyrights"`
	Description        string            `json:"description"`
	HTMLDescription    string            `json:"html_description"`
	Explicit           bool              `json:"explicit"`
	ExternalURLs       map[string]string `json:"external_urls"`
	Href               string            `json:"href"`
	ID                 string            `json:"id"`
	Images             []Image           `json:"images"`
	IsExternallyHosted bool              `json:"is_externally_hosted"`
	Languages          []string          `json:"languages"`
	MediaType          string            `json:"media_type"`
	Name               string            `json:"name"`
	Publisher          string            `json:"publisher"`
	Type               string            `json:"type"`
	URI                string            `json:"uri"`
	TotalEpisodes      int               `json:"total_episodes"`
}

type ShowCopyright struct {
	Text string `json:"text"`
	Type string `json:"type"`
}
