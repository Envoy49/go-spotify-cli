package auth

import (
	"encoding/base64"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"net/url"
	"strings"
)

type TokenResponse struct {
	AccessToken  string `json:"access_token"`
	RefreshToken string `json:"refresh_token"`
	ExpiresIn    uint   `json:"expires_in"`
	TokenType    string `json:"token_type"`
}

func GetAccessToken(clientID, clientSecret, authCode, redirectURI string) (string, uint, error) {
	data := url.Values{}
	data.Set("grant_type", "authorization_code")
	data.Set("code", authCode)
	data.Set("redirect_uri", redirectURI)

	req, err := http.NewRequest("POST", "https://accounts.spotify.com/api/token", strings.NewReader(data.Encode()))
	if err != nil {
		return "", 0, err
	}

	req.Header.Add("Content-Type", "application/x-www-form-urlencoded")
	req.Header.Add("Authorization", "Basic "+base64.StdEncoding.EncodeToString([]byte(clientID+":"+clientSecret)))

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return "", 0, err
	}

	defer func() {
		err := resp.Body.Close()
		if err != nil {
			fmt.Println("Error closing request for /auth", err)
		}
	}()

	body, err := io.ReadAll(resp.Body)

	if err != nil {
		return "", 0, err
	}

	if resp.StatusCode != 200 {
		return "", 0, fmt.Errorf("spotify: got %d status code: %s", resp.StatusCode, body)
	}

	var tokenResponse TokenResponse
	err = json.Unmarshal(body, &tokenResponse)

	if err != nil {
		return "", 0, err
	}

	return tokenResponse.AccessToken, tokenResponse.ExpiresIn, nil
}
