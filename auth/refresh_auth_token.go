package auth

import (
	"encoding/base64"
	"encoding/json"
	"fmt"
	"github.com/sirupsen/logrus"
	"go-spotify-cli/config"
	"io"
	"net/http"
	"net/url"
	"strings"
)

func RefreshAuthToken(refreshToken string) (*FetchTokenResponse, error) {
	data := url.Values{}
	data.Set("grant_type", "refresh_token")
	data.Set("client_id", config.GlobalConfig.ClientId)
	data.Set("refresh_token", refreshToken)

	req, err := http.NewRequest("POST", "https://accounts.spotify.com/api/token", strings.NewReader(data.Encode()))
	if err != nil {
		return nil, err
	}

	req.Header.Add("Content-Type", "application/x-www-form-urlencoded")
	req.Header.Add("Authorization", "Basic "+base64.StdEncoding.EncodeToString([]byte(config.GlobalConfig.ClientId+":"+config.GlobalConfig.ClientSecret)))

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return nil, err
	}

	defer func() {
		err := resp.Body.Close()
		if err != nil {
			logrus.WithError(err).Error("Error closing request for fetch auth token")
		}
	}()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	if resp.StatusCode != 200 {
		return nil, fmt.Errorf("spotify: got %d status code: %s", resp.StatusCode, body)
	}

	var tokenResponse TokenResponse
	err = json.Unmarshal(body, &tokenResponse)

	if err != nil {
		return nil, err
	}

	fetchResponse := &FetchTokenResponse{
		AccessToken: tokenResponse.AccessToken,
		ExpiresIn:   tokenResponse.ExpiresIn,
	}

	return fetchResponse, nil

}
