package config

import (
	"github.com/sirupsen/logrus"
	"go-spotify-cli/constants"
	"gopkg.in/yaml.v2"
	"os"
	"path/filepath"
)

func WriteToHomeDirectory(clientSecret string, clientId string) {
	configData := &EnvVarConfig{
		ClientId:     clientId,
		ClientSecret: clientSecret,
	}

	// Get the home directory for the current user
	homeDir, err := os.UserHomeDir()
	if err != nil {
		logrus.WithError(err).Error("Error getting home directory")
		return
	}

	// Define the folder and file paths
	folderPath := filepath.Join(homeDir, "."+constants.ProjectName)
	filePath := filepath.Join(folderPath, constants.ProjectName+".yaml")

	// Check if the folder already exists, if not then create it
	if _, err := os.Stat(folderPath); os.IsNotExist(err) {
		if err := os.Mkdir(folderPath, 0755); err != nil {
			logrus.WithError(err).Error("Error creating folder")
			return
		}
		logrus.Info("Folder created:", folderPath)
	}

	// Marshal the configData into YAML format
	data, err := yaml.Marshal(configData)
	if err != nil {
		logrus.WithError(err).Error("Error marshalling data to YAML")
		return
	}

	// Write the YAML data to the file
	if err := os.WriteFile(filePath, data, 0644); err != nil {
		logrus.WithError(err).Error("Error writing to file")
		return
	}

	GlobalConfig = Config{
		ClientId:     configData.ClientId,
		ClientSecret: configData.ClientSecret,
	}
	logrus.Info("Configuration saved:", filePath)
}
