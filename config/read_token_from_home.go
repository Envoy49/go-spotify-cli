package config

import (
	"github.com/sirupsen/logrus"
	"go-spotify-cli/constants"
	"gopkg.in/yaml.v2"
	"os"
	"path/filepath"
	"time"
)

func isTokenExpired(expiryTime time.Time) bool {
	return time.Now().After(expiryTime)
}

func ReadTokenFromHome(tokenType string) *TokenStructure {
	// Get the home directory for the current user
	homeDir, err := os.UserHomeDir()
	if err != nil {
		logrus.WithError(err).Error("Error getting home directory")
		return nil
	}

	// Define the folder and file paths
	folderPath := filepath.Join(homeDir, "."+constants.ProjectName)
	filePath := filepath.Join(folderPath, constants.ProjectName+"-env.yaml")

	// Define an instance to store the current file's data
	currentData := &TokenStructure{}

	// Read the existing file (if it exists) and unmarshal its data
	if fileData, err := os.ReadFile(filePath); err == nil {
		if err := yaml.Unmarshal(fileData, &currentData); err != nil {
			logrus.WithError(err).Error("Error unmarshalling existing data from file")
			return nil
		}
	}

	if tokenType == "userReadToken" {
		expiresIn := time.Unix(currentData.UserReadTokenExpiresIn, 0)
		if isTokenExpired(expiresIn) {
			return &TokenStructure{
				UserReadRefreshToken: currentData.UserReadRefreshToken,
			}
		}
		return &TokenStructure{
			UserReadToken: currentData.UserReadToken,
		}
	}

	if tokenType == "userModifyToken" {
		expiresIn := time.Unix(currentData.UserModifyTokenExpiresIn, 0)
		if isTokenExpired(expiresIn) {
			return &TokenStructure{
				UserModifyRefreshToken: currentData.UserModifyRefreshToken,
			}
		}
		return &TokenStructure{
			UserModifyToken: currentData.UserModifyToken,
		}
	}

	return nil
}
