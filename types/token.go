package types

type TokenResponse struct {
	AccessToken  string `json:"access_token"`
	RefreshToken string `json:"refresh_token"`
	ExpiresIn    uint   `json:"expires_in"`
	TokenType    string `json:"token_type"`
}

type FetchTokenResponse struct {
	AccessToken  string
	RefreshToken string
	ExpiresIn    uint
}

type FetchAuthTokenParams struct {
	AuthCode     string
	RedirectURI  string
	RefreshToken string
}
