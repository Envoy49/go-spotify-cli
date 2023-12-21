package config

import (
	"github.com/sirupsen/logrus"
	"go-spotify-cli/constants"
	"gopkg.in/yaml.v2"
	"os"
	"path/filepath"
	"time"
)

var AuthTokenData = make(chan CombinedTokenStructure)

func getTokenExpiryTime(expiresIn int64) time.Time { // expiresIn should be actual time when it is going to expire
	return time.Now().Add(time.Second * time.Duration(expiresIn))
}

func checkModifyToken(current, newToken UserModifyTokenStructure) UserModifyTokenStructure {
	if newToken.UserModifyToken != "" {
		current.UserModifyToken = newToken.UserModifyToken
	}
	if newToken.UserModifyRefreshToken != "" {
		current.UserModifyRefreshToken = newToken.UserModifyRefreshToken
	}
	if newToken.UserModifyTokenExpiresIn != 0 {
		current.UserModifyTokenExpiresIn = getTokenExpiryTime(newToken.UserModifyTokenExpiresIn).Unix()
	}
	return current
}

func checkReadToken(current, newToken UserReadTokenStructure) UserReadTokenStructure {
	if newToken.UserReadToken != "" {
		current.UserReadToken = newToken.UserReadToken
	}
	if newToken.UserReadRefreshToken != "" {
		current.UserReadRefreshToken = newToken.UserReadRefreshToken
	}
	if newToken.UserReadTokenExpiresIn != 0 {
		current.UserReadTokenExpiresIn = getTokenExpiryTime(newToken.UserReadTokenExpiresIn).Unix()
	}
	return current
}

func WriteTokenToHomeDirectory(configData *CombinedTokenStructure, initiateChannel bool) {
	// Get the home directory for the current user
	homeDir, err := os.UserHomeDir()
	if err != nil {
		logrus.WithError(err).Error("Error getting home directory")
		return
	}

	// Define the folder and file paths
	folderPath := filepath.Join(homeDir, "."+constants.ProjectName)
	filePath := filepath.Join(folderPath, constants.ProjectName+"-env.yaml")

	// Check if the folder already exists, if not then create it
	if _, err := os.Stat(folderPath); os.IsNotExist(err) {
		if err := os.Mkdir(folderPath, 0755); err != nil {
			logrus.WithError(err).Error("Error creating folder")
			return
		}
		logrus.Println("Folder created:", folderPath)
	}
	// Define an instance to store the current file's data
	currentData := CombinedTokenStructure{}

	// Read the existing file (if it exists) and unmarshal its data
	if fileData, err := os.ReadFile(filePath); err == nil {
		if err := yaml.Unmarshal(fileData, &currentData); err != nil {
			logrus.WithError(err).Error("Error unmarshalling existing data from file")
			return
		}
	}
	if configData != nil {
		currentData.ModifyToken = checkModifyToken(currentData.ModifyToken, configData.ModifyToken)
		currentData.ReadToken = checkReadToken(currentData.ReadToken, configData.ReadToken)
	}

	if configData == nil {
		// Flush all tokens if configData not provided
		if err := os.WriteFile(filePath, nil, 0644); err != nil {
			logrus.WithError(err).Error("Error writing to file")
			return
		}
	} else {
		// Marshal the updated data
		data, err := yaml.Marshal(currentData)
		if err != nil {
			logrus.WithError(err).Error("Error marshalling data to YAML")
			return
		}

		// Write the updated YAML data to the file
		if err := os.WriteFile(filePath, data, 0644); err != nil {
			logrus.WithError(err).Error("Error writing to file")
			return
		}

		if initiateChannel {
			AuthTokenData <- currentData
		}

		logrus.Println("Token information saved to:", filePath)
	}

}
