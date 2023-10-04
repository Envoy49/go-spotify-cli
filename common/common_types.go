package common

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
