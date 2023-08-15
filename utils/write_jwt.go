package utils

import (
	"fmt"
	"go-spotify-cli/constants"
	"log"
	"os"
)

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

	data := map[string]interface{}{
		"jwtToken":  token,
		"expiresIn": expiresIn,
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

	fmt.Println("---------> New token is written to cache")

	return nil
}
