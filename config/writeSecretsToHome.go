package config

import (
	"os"
	"path/filepath"

	"github.com/sirupsen/logrus"
	"gopkg.in/yaml.v2"
)

func WriteSecretsToHomeDirectory(cfg *Config) (*Config, error) { // pass this a struct one parameter
	var configData *EnvVarConfig

	if cfg != nil {
		configData = &EnvVarConfig{
			ClientId:     cfg.ClientId,
			ClientSecret: cfg.ClientSecret,
		}
	}

	// Get the home directory for the current user
	homeDir, err := os.UserHomeDir()
	if err != nil {
		logrus.WithError(err).Error("Error getting home directory")
		return nil, err
	}

	// Define the folder and file paths
	folderPath := filepath.Join(homeDir, "."+projectName)
	filePath := filepath.Join(folderPath, projectName+".yaml")

	// Check if the folder already exists, if not then create it
	if _, err := os.Stat(folderPath); os.IsNotExist(err) {
		if err := os.Mkdir(folderPath, 0755); err != nil {
			logrus.WithError(err).Error("Error creating folder")
			return nil, err
		}
		logrus.Println("Folder created:", folderPath)
	}

	// Marshal the configData into YAML format
	data, err := yaml.Marshal(configData)
	if err != nil {
		logrus.WithError(err).Error("Error marshaling data to YAML")
		return nil, err
	}

	// Write the YAML data to the file
	if err := os.WriteFile(filePath, data, 0644); err != nil {
		logrus.WithError(err).Error("Error writing to file")
		return nil, err
	}

	if cfg != nil {
		logrus.Println("Configuration saved:", filePath)
	} else {
		logrus.Println("All saved secrets deleted")
		return nil, nil
	}

	return &Config{
		ClientId:     configData.ClientId,
		ClientSecret: configData.ClientSecret,
	}, nil
}
