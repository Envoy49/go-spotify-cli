package utils

import (
	"bufio"
	"fmt"
	"go-spotify-cli/constants"
	"os"
	"strings"
)

func ReadJWTToken() (string, error) {
	file, err := os.OpenFile(constants.TempFileName, os.O_RDWR|os.O_CREATE, 0644)
	if err != nil {
		return "", fmt.Errorf("error opening or creating file: %w", err)
	}

	defer func() {
		closeErr := file.Close()
		if closeErr != nil {
			fmt.Println("Error closing file:", closeErr)
		}
	}()

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()
		parts := strings.SplitN(line, "=", 2)
		if len(parts) == 2 && parts[0] == "jwtToken" {
			fmt.Println("---------> Token Found")
			return parts[1], nil // Found the token
		}
	}

	if err := scanner.Err(); err != nil {
		return "", fmt.Errorf("error reading file: %w", err)
	}

	fmt.Println("---------> Token Not Found")

	return "", nil // Token not found
}
