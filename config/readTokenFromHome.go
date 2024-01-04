package config

import (
	"github.com/envoy49/go-spotify-cli/constants"
	"github.com/envoy49/go-spotify-cli/types"
	"github.com/sirupsen/logrus"
	"gopkg.in/yaml.v2"
	"os"
	"path/filepath"
	"time"
)

func isTokenExpired(expiryTime time.Time) bool {
	return time.Now().After(expiryTime)
}

func ReadTokenFromHome(tokenType constants.TokenType) *types.CombinedTokenStructure {
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
	currentData := &types.CombinedTokenStructure{}

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
			return &types.CombinedTokenStructure{
				ReadToken: types.UserReadTokenStructure{
					UserReadRefreshToken: currentData.ReadToken.UserReadRefreshToken,
				},
			}
		}

		return &types.CombinedTokenStructure{
			ReadToken: types.UserReadTokenStructure{
				UserReadToken: currentData.ReadToken.UserReadToken,
			},
		}
	}

	if tokenType == constants.ModifyToken {
		expiresIn := time.Unix(currentData.ModifyToken.UserModifyTokenExpiresIn, 0)
		if isTokenExpired(expiresIn) {
			return &types.CombinedTokenStructure{
				ModifyToken: types.UserModifyTokenStructure{
					UserModifyRefreshToken: currentData.ModifyToken.UserModifyRefreshToken,
				},
			}
		}
		return &types.CombinedTokenStructure{
			ModifyToken: types.UserModifyTokenStructure{
				UserModifyToken: currentData.ModifyToken.UserModifyToken,
			},
		}
	}

	if tokenType == constants.LibraryRead {
		expiresIn := time.Unix(currentData.LibraryReadToken.UserLibraryReadTokenExpiresIn, 0)
		if isTokenExpired(expiresIn) {
			return &types.CombinedTokenStructure{
				LibraryReadToken: types.UserLibraryReadTokenStructure{
					UserLibraryReadRefreshToken: currentData.LibraryReadToken.UserLibraryReadRefreshToken,
				},
			}
		}
		return &types.CombinedTokenStructure{
			LibraryReadToken: types.UserLibraryReadTokenStructure{
				UserLibraryReadToken: currentData.LibraryReadToken.UserLibraryReadToken,
			},
		}
	}

	return nil
}
