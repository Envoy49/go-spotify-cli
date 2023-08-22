package utils

import (
	"fmt"
	"go-spotify-cli/constants"
	"log"
	"os"
	"time"
)

func getTokenExpiryTime(expiresIn uint) time.Time { // expires in should be actual time when it is going to expire
	return time.Now().Add(time.Second * time.Duration(expiresIn))
}

func WriteJWTToken(token string, expiresIn uint) error {
	file, err := os.OpenFile(constants.TempFileName, os.O_RDWR|os.O_CREATE, 0644)
	if err != nil {
		return fmt.Errorf("error opening or creating file: %w", err)
	}
	defer func() {
		if closeErr := file.Close(); closeErr != nil {
			// Log the close error
			fmt.Println("Error closing file:", closeErr)
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
			log.Printf("Unsupported type for key %s\n", key)
		}
	}

	fmt.Println("Token cache miss")

	return nil
}
