package config

import (
	"os"
	"path/filepath"

	"github.com/envoy49/go-spotify-cli/types"
	"github.com/sirupsen/logrus"
	"gopkg.in/yaml.v2"
)

const (
	projectName = "go-spotify-cli"
)

var GlobalConfig types.Config

func LoadConfiguration() {
	homeDir, err := os.UserHomeDir()
	if err != nil {
		logrus.WithError(err).Error("Error getting home directory")
		return
	}

	folderPath := filepath.Join(homeDir, "."+projectName)
	filePath := filepath.Join(folderPath, projectName+".yaml")

	data, err := os.ReadFile(filePath)
	if err != nil {
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

func VerifyConfigExists() bool {
	// Get the home directory for the current user
	homeDir, err := os.UserHomeDir()
	if err != nil {
		return false
	}

	// Define the folder and file paths
	folderPath := filepath.Join(homeDir, "."+projectName)
	filePath := filepath.Join(folderPath, projectName+".yaml")

	// Check if the folder exists
	folderExists, err := os.Stat(folderPath)
	if err != nil || os.IsNotExist(err) || !folderExists.IsDir() {
		return false
	}

	// Check if the file exists
	_, err = os.Stat(filePath)

	if err != nil {
		return false
	}

	if len(GlobalConfig.ClientSecret) == 0 || len(GlobalConfig.ClientId) == 0 {
		return false
	}

	return true
}
