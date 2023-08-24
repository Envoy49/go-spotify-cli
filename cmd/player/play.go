package player

import (
	"fmt"
	"go-spotify-cli/utils"
	"io"
	"net/http"
)

func Play(accessToken string) error {
	req, err := http.NewRequest("PUT", "https://api.spotify.com/v1/me/player/play", nil)
	if err != nil {
		return err
	}

	req.Header.Set("Authorization", "Bearer "+accessToken)
	req.Header.Set("Content-Type", "application/json")
	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		return err
	}

	defer func() {
		err := resp.Body.Close()
		if err != nil {
			utils.PrintError("Error closing request for /auth", err)
		}
	}()

	if resp.StatusCode != http.StatusNoContent {
		body, _ := io.ReadAll(resp.Body)
		return fmt.Errorf("could not play: %v", string(body))
	}

	return nil
}
