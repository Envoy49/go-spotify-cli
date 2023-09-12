package utils

import (
	"fmt"
	"github.com/sirupsen/logrus"
	"go-spotify-cli/constants"
	"log"
	"os"
	"path/filepath"
	"time"
)

func getTokenExpiryTime(expiresIn uint) time.Time { // expires in should be actual time when it is going to expire
	return time.Now().Add(time.Second * time.Duration(expiresIn))
}

var AuthToken = make(chan string)

func WriteJWTToken(token string, expiresIn uint) error {
	tempDir := os.TempDir()
	fullTempFilePath := filepath.Join(tempDir, constants.TempFileName)
	file, err := os.OpenFile(fullTempFilePath, os.O_RDWR|os.O_CREATE, 0644)
	if err != nil {
		return fmt.Errorf("error opening or creating file: %w", err)
	}
	defer func() {
		if closeErr := file.Close(); closeErr != nil {
			logrus.WithError(closeErr).Error("Error closing file")
		}
	}()

	expiryTime := getTokenExpiryTime(expiresIn)
	expiryTimeString := expiryTime.Format(time.RFC3339)

	data := map[string]interface{}{
		"jwtToken":  token,
		"expiresIn": expiryTimeString,
	}

	for key, value := range data {
		switch v := value.(type) {
		case string:
			_, err := fmt.Fprintf(file, "%s=%s\n", key, v)
			if err != nil {
				log.Fatal(err)
			}
		case uint:
			_, err := fmt.Fprintf(file, "%s=%d\n", key, v)
			if err != nil {
				log.Fatal(err)
			}
		default:
			logrus.Warn("Unsupported type for key %s\n", key)
		}
	}

	logrus.Info("Token cache miss")
	// After writing token to the cache pass it to the command, so it can continue
	AuthToken <- token

	return nil
}
