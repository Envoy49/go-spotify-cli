package config

import (
	"github.com/envoy49/go-spotify-cli/commands/commandTypes"
	"os"
	"path/filepath"
	"time"

	"github.com/envoy49/go-spotify-cli/constants"
	"github.com/sirupsen/logrus"
	"gopkg.in/yaml.v2"
)

func isTokenExpired(expiryTime time.Time) bool {
	return time.Now().After(expiryTime)
}

func ReadTokenFromHome(tokenType constants.TokenType) *commandTypes.CombinedTokenStructure {
	// Get the home directory for the current user
	homeDir, err := os.UserHomeDir()
	if err != nil {
		logrus.WithError(err).Error("Error getting home directory")
		return nil
	}

	// Define the folder and file paths
	folderPath := filepath.Join(homeDir, "."+projectName)
	filePath := filepath.Join(folderPath, projectName+"-env.yaml")

	// Define an instance to store the current file's data
	currentData := &commandTypes.CombinedTokenStructure{}

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
			return &commandTypes.CombinedTokenStructure{
				ReadToken: commandTypes.UserReadTokenStructure{
					UserReadRefreshToken: currentData.ReadToken.UserReadRefreshToken,
				},
			}
		}

		return &commandTypes.CombinedTokenStructure{
			ReadToken: commandTypes.UserReadTokenStructure{
				UserReadToken: currentData.ReadToken.UserReadToken,
			},
		}
	}

	if tokenType == constants.ModifyToken {
		expiresIn := time.Unix(currentData.ModifyToken.UserModifyTokenExpiresIn, 0)
		if isTokenExpired(expiresIn) {
			return &commandTypes.CombinedTokenStructure{
				ModifyToken: commandTypes.UserModifyTokenStructure{
					UserModifyRefreshToken: currentData.ModifyToken.UserModifyRefreshToken,
				},
			}
		}
		return &commandTypes.CombinedTokenStructure{
			ModifyToken: commandTypes.UserModifyTokenStructure{
				UserModifyToken: currentData.ModifyToken.UserModifyToken,
			},
		}
	}

	if tokenType == constants.LibraryRead {
		expiresIn := time.Unix(currentData.LibraryReadToken.UserLibraryReadTokenExpiresIn, 0)
		if isTokenExpired(expiresIn) {
			return &commandTypes.CombinedTokenStructure{
				LibraryReadToken: commandTypes.UserLibraryReadTokenStructure{
					UserLibraryReadRefreshToken: currentData.LibraryReadToken.UserLibraryReadRefreshToken,
				},
			}
		}
		return &commandTypes.CombinedTokenStructure{
			LibraryReadToken: commandTypes.UserLibraryReadTokenStructure{
				UserLibraryReadToken: currentData.LibraryReadToken.UserLibraryReadToken,
			},
		}
	}

	return nil
}
