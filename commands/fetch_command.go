package commands

import (
	"encoding/json"
	"fmt"
	"github.com/sirupsen/logrus"
	"go-spotify-cli/common"
	"io"
	"net/http"
	"sync"
)

type PlayerParams struct {
	AccessToken string
	Method      string
	Endpoint    string
	Body        io.Reader
}

//	func printProgress(current, total int) {
//		// Base styles
//		//baseStyle := lipgloss.NewStyle().Foreground(lipgloss.Color("#51e2f5")).Bold(true)
//		percentageStyle := lipgloss.NewStyle().Foreground(lipgloss.AdaptiveColor{Light: "#333", Dark: "#DDD"}).Bold(true)
//		backgroundStyle := lipgloss.NewStyle().Background(lipgloss.Color("#51e2f5"))
//
//		percentage := float64(current) / float64(total)
//		blocksTotal := 50
//
//		// Linear interpolation formula: (end - start) * percentage + start
//		blocksFilled := int(float64(blocksTotal-7)*percentage + 7.0)
//		blocksEmpty := blocksTotal - blocksFilled
//
//		// Styling the filled and empty parts of the bar
//		filledBlocks := backgroundStyle.Render(strings.Repeat(" ", blocksFilled))
//		emptyBlocks := strings.Repeat(" ", blocksEmpty)
//
//		progress := fmt.Sprintf(
//			"%s %3.0f%%",
//			filledBlocks+emptyBlocks,
//			percentage*100,
//		)
//		fmt.Printf("\r%s", percentageStyle.Render(progress))
//	}
var fetchMutex sync.Mutex

func FetchCommand(playerParams *PlayerParams) ([]byte, error) {
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
		var errorDetails common.SpotifyError
		if jsonErr := json.Unmarshal(body, &errorDetails); jsonErr != nil {
			return nil, fmt.Errorf("unexpected error format from Spotify API: %s", string(body))
		}
		return nil, common.SpotifyAPIError{Detail: errorDetails}
	}

	return body, nil
}
