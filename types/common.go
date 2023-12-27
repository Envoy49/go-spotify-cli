package types

type Image struct {
	URL    string `json:"url"`
	Height int    `json:"height"`
	Width  int    `json:"width"`
}

type SpotifyError struct {
	Error struct {
		Status  int    `json:"status"`
		Message string `json:"message"`
		Reason  string `json:"reason"`
	} `json:"error"`
}

type SpotifyAPIError struct {
	Detail SpotifyError
}

func (e SpotifyAPIError) Error() string {
	return e.Detail.Error.Message
}

type UrlParams struct {
	ClientID        string
	RedirectURI     string
	RequestedScopes string
}

type Body struct {
	Href     string `json:"href"`
	Limit    int    `json:"limit"`
	Next     string `json:"next"`
	Offset   int    `json:"offset"`
	Previous string `json:"previous"`
	Total    int    `json:"total"`
}

type SelectionPromptConfig struct {
	Label         string
	FormattedInfo []string
}
