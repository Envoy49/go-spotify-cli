package config

import (
	"encoding/json"
	"github.com/sirupsen/logrus"
	"os"
)

type Config struct {
	ServerUrl       string
	Port            string
	RequestedScopes string
	ClientId        string
	ClientSecret    string
}

var GlobalConfig Config

func LoadConfiguration() {
	file, err := os.Open("./config.json")
	if err != nil {
		logrus.WithError(err).Error("Error opening config.json")
		return
	}
	defer func() {
		err := file.Close()
		if err != nil {
			logrus.WithError(err).Error("Error closing config.json file")
		}
	}()

	decoder := json.NewDecoder(file)
	config := Config{}
	err = decoder.Decode(&config)
	if err != nil {
		logrus.WithError(err).Error("error decoding config")
		return
	}
	GlobalConfig = Config{
		ServerUrl:    config.ServerUrl,
		Port:         config.Port,
		ClientId:     config.ClientId,
		ClientSecret: config.ClientSecret,
	}
}
