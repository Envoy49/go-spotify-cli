package player

import (
	"bytes"
	"encoding/json"

	"github.com/envoy49/go-spotify-cli/config"

	"github.com/envoy49/go-spotify-cli/commands"
	"github.com/envoy49/go-spotify-cli/server"
	"github.com/sirupsen/logrus"
)

func ActivateDevice(cfg *config.Config, deviceIds string) {
	jsonData := map[string]interface{}{
		"device_ids": []string{deviceIds},
		"play":       true,
	}

	jsonDataBytes, err1 := json.Marshal(jsonData)

	if err1 != nil {
		logrus.WithError(err1).Info("error Marshaling Json data in Activate Device")
		return
	}
	requestBody := bytes.NewBuffer(jsonDataBytes)

	token := server.ReadUserModifyTokenOrFetchFromServer(cfg)

	params := &commands.PlayerParams{
		AccessToken: token,
		Method:      "PUT",
		Body:        requestBody,
		Endpoint:    spotifyPlayerEndpoint + "/player",
	}

	_, err := commands.Fetch(params)

	if err != nil {
		logrus.WithError(err).Error("Error activating device")
		return
	} else {
		logrus.Println("Device has been activated")
	}
}
