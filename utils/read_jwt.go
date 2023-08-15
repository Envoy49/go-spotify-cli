package utils

import (
	"bufio"
	"errors"
	"fmt"
	"go-spotify-cli/constants"
	"log"
	"os"
	"strconv"
	"strings"
	"time"
)

func getTokenExpiryTime(expiresIn string) time.Time {
	value, err := strconv.ParseInt(expiresIn, 10, 64)
	if err != nil {
		log.Fatalf("Failed to parse expiredIn: %v", err)
	}

	return time.Now().Add(time.Second * time.Duration(value))
}

func isTokenExpired(expiryTime time.Time) bool {
	return time.Now().After(expiryTime)
}

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
	var token string
	var expiresIn string

	for scanner.Scan() {
		line := scanner.Text()
		parts := strings.SplitN(line, "=", 2)
		if len(parts) != 2 {
			continue
		}

		switch parts[0] {
		case "jwtToken":
			token = parts[1]
		case "expiresIn":
			expiresIn = parts[1]
		}
	}
	fmt.Println("------------->expiresIn", expiresIn)
	tokenExpired := isTokenExpired(getTokenExpiryTime(expiresIn))

	if tokenExpired {
		return "", errors.New("TOKEN EXPIRED")
	}

	if err := scanner.Err(); err != nil {
		return "", fmt.Errorf("error reading file: %w", err)
	}

	if token == "" {
		return "", errors.New("token not found")
	}

	return token, nil // Token not found
}
