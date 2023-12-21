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

func ReadTokenFromHome(tokenType constants.TokenType) *CombinedTokenStructure {
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
	currentData := &CombinedTokenStructure{}

	// Read the existing file (if it exists) and unmarshal its data
	if fileData, err := os.ReadFile(filePath); err == nil {
		if err := yaml.Unmarshal(fileData, &currentData); err != nil {
			logrus.WithError(err).Error("Error unmarshalling existing data from file")
			return nil
		}
	}

	if tokenType == constants.ReadToken {
		expiresIn := time.Unix(currentData.ReadToken.UserReadTokenExpiresIn, 0)
		if isTokenExpired(expiresIn) {
			return &CombinedTokenStructure{
				ReadToken: UserReadTokenStructure{
					UserReadRefreshToken: currentData.ReadToken.UserReadRefreshToken,
				},
			}
		}

		return &CombinedTokenStructure{
			ReadToken: UserReadTokenStructure{
				UserReadToken: currentData.ReadToken.UserReadToken,
			},
		}
	}

	if tokenType == constants.ModifyToken {
		expiresIn := time.Unix(currentData.ModifyToken.UserModifyTokenExpiresIn, 0)
		if isTokenExpired(expiresIn) {
			return &CombinedTokenStructure{
				ModifyToken: UserModifyTokenStructure{
					UserModifyRefreshToken: currentData.ModifyToken.UserModifyRefreshToken,
				},
			}
		}
		return &CombinedTokenStructure{
			ModifyToken: UserModifyTokenStructure{
				UserModifyToken: currentData.ModifyToken.UserModifyToken,
			},
		}
	}

	return nil
}
