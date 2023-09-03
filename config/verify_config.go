package config

import (
	"go-spotify-cli/constants"
	"os"
	"path/filepath"
)

func VerifyConfigExists() bool {
	// Get the home directory for the current user
	homeDir, err := os.UserHomeDir()
	if err != nil {
		return false
	}

	// Define the folder and file paths
	folderPath := filepath.Join(homeDir, "."+constants.ProjectName)
	filePath := filepath.Join(folderPath, constants.ProjectName+".yaml")

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
