package config

import (
	"github.com/sirupsen/logrus"
	"go-spotify-cli/constants"
	"go-spotify-cli/types"
	"gopkg.in/yaml.v2"
	"os"
	"path/filepath"
)

var GlobalConfig types.Config

func LoadConfiguration() {
	homeDir, err := os.UserHomeDir()
	if err != nil {
		logrus.WithError(err).Error("Error getting home directory")
		return
	}

	folderPath := filepath.Join(homeDir, "."+constants.ProjectName)
	filePath := filepath.Join(folderPath, constants.ProjectName+".yaml")

	data, err := os.ReadFile(filePath)
	if err != nil {
		logrus.WithError(err).Error("Error reading the file")
		return
	}

	var config types.EnvVarConfig
	// Unmarshal the YAML data into a Configuration struct
	err = yaml.Unmarshal(data, &config)
	if err != nil {
		logrus.WithError(err).Error("Error unmarshalling YAML data")
		return
	}

	GlobalConfig = types.Config{
		ClientId:     config.ClientId,
		ClientSecret: config.ClientSecret,
	}
}
