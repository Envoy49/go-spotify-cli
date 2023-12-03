package config

import (
	"github.com/sirupsen/logrus"
	"go-spotify-cli/constants"
	"gopkg.in/yaml.v2"
	"os"
	"path/filepath"
	"time"
)

type TokenStructure struct {
	UserModifyToken          string `yaml:"UserModifyToken"`
	UserModifyRefreshToken   string `yaml:"UserModifyRefreshToken"`
	UserModifyTokenExpiresIn int64  `yaml:"UserModifyTokenExpiresIn"`
	UserReadToken            string `yaml:"UserReadToken"`
	UserReadRefreshToken     string `yaml:"UserReadRefreshToken"`
	UserReadTokenExpiresIn   int64  `yaml:"UserReadTokenExpiresIn"`
}

func getTokenExpiryTime(expiresIn int64) time.Time { // expires in should be actual time when it is going to expire
	return time.Now().Add(time.Second * time.Duration(expiresIn))
}

var AuthTokenData = make(chan TokenStructure)

func WriteTokenToHomeDirectory(configData *TokenStructure, initiateChannel bool) {
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
	currentData := TokenStructure{}

	// Read the existing file (if it exists) and unmarshal its data
	if fileData, err := os.ReadFile(filePath); err == nil {
		if err := yaml.Unmarshal(fileData, &currentData); err != nil {
			logrus.WithError(err).Error("Error unmarshalling existing data from file")
			return
		}
	}
	// Update the fields of currentData with the non-empty fields of configData
	if configData.UserModifyToken != "" {
		currentData.UserModifyToken = configData.UserModifyToken
	}

	if configData.UserModifyRefreshToken != "" {
		currentData.UserModifyRefreshToken = configData.UserModifyRefreshToken
	}

	if configData.UserModifyTokenExpiresIn != 0 {
		currentData.UserModifyTokenExpiresIn = getTokenExpiryTime(configData.UserModifyTokenExpiresIn).Unix()
	}

	if configData.UserReadToken != "" {
		currentData.UserReadToken = configData.UserReadToken
	}

	if configData.UserReadRefreshToken != "" {
		currentData.UserReadRefreshToken = configData.UserReadRefreshToken
	}

	if configData.UserReadTokenExpiresIn != 0 {
		currentData.UserReadTokenExpiresIn = getTokenExpiryTime(configData.UserReadTokenExpiresIn).Unix()
	}

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
