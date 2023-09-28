package cmd

import (
	"encoding/json"
	"fmt"
	"github.com/sirupsen/logrus"
	"go-spotify-cli/common"
	"io"
	"net/http"
)

type PlayerParams struct {
	AccessToken string
	Method      string
	Endpoint    string
	Body        io.Reader
}

func FetchCommand(playerParams *PlayerParams) ([]byte, error) {
	req, err := http.NewRequest(
		playerParams.Method,
		playerParams.Endpoint,
		playerParams.Body,
	)

	if err != nil {
		return nil, err
	}

	req.Header.Set("Authorization", "Bearer "+playerParams.AccessToken)
	req.Header.Set("Content-Type", "application/json")
	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		return nil, err
	}

	body, readErr := io.ReadAll(resp.Body)
	if readErr != nil {
		return nil, readErr
	}

	defer func() {
		if resp != nil && resp.Body != nil {
			if bodyErr := resp.Body.Close(); bodyErr != nil {
				logrus.WithError(bodyErr).Error("Error closing response body")
			}
		}
	}()

	if resp.StatusCode != http.StatusOK && resp.StatusCode != http.StatusNoContent {
		var errorDetails common.SpotifyError
		if jsonErr := json.Unmarshal(body, &errorDetails); jsonErr != nil {
			return nil, fmt.Errorf("unexpected error format from Spotify API: %s", string(body))
		}
		return nil, common.SpotifyAPIError{Detail: errorDetails}
	}

	return body, nil
}
