package commands

import (
	"fmt"
	"io"
	"net/http"
)

type PlayerParams struct {
	AccessToken string
	Method      string
	Endpoint    string
}

func Player(playerParams *PlayerParams) error {
	req, err := http.NewRequest(playerParams.Method, "https://api.spotify.com/v1/me/player"+fmt.Sprint(playerParams.Endpoint), nil)
	if err != nil {
		return err
	}

	req.Header.Set("Authorization", "Bearer "+playerParams.AccessToken)
	req.Header.Set("Content-Type", "application/json")
	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		return err
	}

	defer func() {
		if bodyErr := resp.Body.Close(); bodyErr != nil {
			// log the error or handle it in another appropriate manner
			fmt.Println("Error closing response body:", bodyErr)
		}
	}()

	if resp.StatusCode != http.StatusNoContent {
		body, _ := io.ReadAll(resp.Body)
		return fmt.Errorf(string(body))
	}

	return nil
}
