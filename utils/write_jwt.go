package utils

import (
	"fmt"
	"go-spotify-cli/constants"
	"os"
)

func WriteJWTToken(token string) error {
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

	_, err = fmt.Fprintf(file, "jwtToken=%s\n", token)
	if err != nil {
		return fmt.Errorf("error writing to file: %w", err)
	}

	fmt.Println("---------> New token is written to cache")

	return nil
}
