package commands

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"sync"

	"github.com/envoy49/go-spotify-cli/commands/cmdTypes"
	"github.com/sirupsen/logrus"
)

type PlayerParams struct {
	AccessToken string
	Method      string
	Endpoint    string
	Body        io.Reader
}

var fetchMutex sync.Mutex

func Fetch(playerParams *PlayerParams) ([]byte, error) {
	fetchMutex.Lock()
	defer fetchMutex.Unlock()

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
		var errorDetails cmdTypes.SpotifyError
		if jsonErr := json.Unmarshal(body, &errorDetails); jsonErr != nil {
			return nil, fmt.Errorf("unexpected error format from Spotify API: %s", string(body))
		}
		return nil, cmdTypes.SpotifyAPIError{Detail: errorDetails}
	}

	return body, nil
}
