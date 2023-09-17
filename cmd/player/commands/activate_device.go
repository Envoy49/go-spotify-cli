package commands

import (
	"bytes"
	"encoding/json"
	"github.com/sirupsen/logrus"
	"go-spotify-cli/cmd/player"
	"go-spotify-cli/server"
)

func ActivateDevice(deviceIds string) {
	jsonData := map[string]interface{}{
		"device_ids": []string{deviceIds},
		"play":       true,
	}

	jsonDataBytes, err1 := json.Marshal(jsonData)

	if err1 != nil {
		logrus.WithError(err1).Info("the reason")
	}
	requestBody := bytes.NewBuffer(jsonDataBytes)

	token := server.ReadUserModifyTokenOrFetchFromServer()

	params := &commands.PlayerParams{
		AccessToken: token,
		Method:      "PUT",
		Body:        requestBody,
	}

	_, err := commands.FetchCommand(params)

	if err != nil {
		logrus.WithError(err).Error("Error activating device")
	} else {
		logrus.Info("Device has been activated")
	}
}
