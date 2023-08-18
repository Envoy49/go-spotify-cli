package utils

import (
	"bufio"
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
		log.Println("Failed to parse expiredIn", err)
	}

	return time.Now().Add(time.Second * time.Duration(value))
}

func isTokenExpired(expiryTime time.Time) bool {
	return time.Now().After(expiryTime)
}

func ReadJWTToken() string {
	file, err := os.OpenFile(constants.TempFileName, os.O_RDWR|os.O_CREATE, 0644)
	if err != nil {
		return ""
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

	if token == "" || expiresIn == "" {
		return ""
	}

	tokenExpired := isTokenExpired(getTokenExpiryTime(expiresIn))

	if tokenExpired {
		return ""
	}

	if err := scanner.Err(); err != nil {
		return ""
	}

	fmt.Println("Token cache hit")

	return token
}
