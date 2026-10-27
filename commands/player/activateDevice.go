package player

import (
	"bytes"
	"encoding/json"

	"github.com/envoy49/go-spotify-cli/config"

	"github.com/envoy49/go-spotify-cli/commands"
	"github.com/envoy49/go-spotify-cli/server"
	"github.com/sirupsen/logrus"
)

func transferPlayback(accessToken, deviceID string, startPlayback bool) error {
	jsonData := struct {
		DeviceIDs []string `json:"device_ids"`
		Play      bool     `json:"play"`
	}{
		DeviceIDs: []string{deviceID},
		Play:      startPlayback,
	}

	jsonDataBytes, err := json.Marshal(jsonData)
	if err != nil {
		return err
	}
	requestBody := bytes.NewBuffer(jsonDataBytes)
	params := &commands.PlayerParams{
		AccessToken: accessToken,
		Method:      "PUT",
		Body:        requestBody,
		Endpoint:    spotifyPlayerEndpoint + "/player",
	}

	_, err = commands.Fetch(params)
	return err
}

func ActivateDevice(cfg *config.Config, deviceID string) {
	token := server.ReadUserModifyTokenOrFetchFromServer(cfg)
	err := transferPlayback(token, deviceID, true)
	if err != nil {
		logrus.WithError(err).Error("Error activating device")
		return
	}

	logrus.Println("Device has been activated")
}
