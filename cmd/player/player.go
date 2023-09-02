package commands

import (
	"fmt"
	"github.com/sirupsen/logrus"
	"io"
	"net/http"
)

type PlayerParams struct {
	AccessToken string
	Method      string
	Endpoint    string
}

func Player(playerParams *PlayerParams) (*http.Response, string, error) {
	req, err := http.NewRequest(
		playerParams.Method,
		"https://api.spotify.com/v1/me/player"+playerParams.Endpoint,
		nil,
	)

	if err != nil {
		return nil, "", err
	}

	req.Header.Set("Authorization", "Bearer "+playerParams.AccessToken)
	req.Header.Set("Content-Type", "application/json")
	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		return nil, "", err
	}

	body, readErr := io.ReadAll(resp.Body)
	if readErr != nil {
		return nil, "", readErr
	}

	defer func() {
		if resp != nil && resp.Body != nil {
			if bodyErr := resp.Body.Close(); bodyErr != nil {
				// log the error or handle it in another appropriate manner
				logrus.WithError(bodyErr).Error("Error closing response body")
			}
		}
	}()

	if resp.StatusCode != http.StatusOK && resp.StatusCode != http.StatusNoContent {
		return nil, "", fmt.Errorf(string(body))
	}

	return resp, string(body), nil
}
