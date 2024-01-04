package auth

import (
	"encoding/base64"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"net/url"
	"strings"

	"github.com/envoy49/go-spotify-cli/config"
	"github.com/envoy49/go-spotify-cli/types"
	"github.com/sirupsen/logrus"
)

func setAuthTokenQueryParams(authCode string, redirectURI string) url.Values {
	data := url.Values{}
	data.Set("grant_type", "authorization_code")
	data.Set("code", authCode)
	data.Set("redirect_uri", redirectURI)
	return data
}

func setRefreshTokenQueryParams(refreshToken string) url.Values {
	data := url.Values{}
	data.Set("grant_type", "refresh_token")
	data.Set("client_id", config.GlobalConfig.ClientId)
	data.Set("refresh_token", refreshToken)
	return data
}

func FetchAuthToken(params *types.FetchAuthTokenParams) (*types.FetchTokenResponse, error) {
	var data url.Values

	if len(params.RefreshToken) > 0 {
		data = setRefreshTokenQueryParams(params.RefreshToken)
	} else {
		data = setAuthTokenQueryParams(params.AuthCode, params.RedirectURI)
	}

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

	var tokenResponse types.TokenResponse
	err = json.Unmarshal(body, &tokenResponse)

	if err != nil {
		return nil, err
	}

	fetchResponse := &types.FetchTokenResponse{
		AccessToken:  tokenResponse.AccessToken,
		RefreshToken: tokenResponse.RefreshToken,
		ExpiresIn:    tokenResponse.ExpiresIn,
	}

	return fetchResponse, nil
}
