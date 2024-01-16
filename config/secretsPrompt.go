package config

import (
	"bufio"
	"fmt"
	"os"
	"regexp"
	"strings"

	"github.com/envoy49/go-spotify-cli/loader"
)

func SecretsPrompt(cfg *Config) *Config {
	if !IsEmptyConfig(cfg) {
		return nil
	}

	if VerifyConfigExists(cfg) {
		return nil
	}

	loader.Stop()

	SecretsSetupPrompt()

	// Function to validate Spotify Client ID and Client Secret
	validate := func(input string) bool {
		var clientIDRegex = regexp.MustCompile(`^[a-zA-Z0-9]{32}$`)
		return clientIDRegex.MatchString(input)
	}

	reader := bufio.NewReader(os.Stdin)

	// Function to prompt and validate input
	promptInput := func(promptText string, validator func(string) bool) string {
		for {
			fmt.Print(promptText)
			input, _ := reader.ReadString('\n')
			input = strings.TrimSpace(input)

			if validator(input) {
				fmt.Println("✅")
				return input
			} else {
				fmt.Println("❌ Invalid input. Please try again.")
			}
		}
	}

	// Get and validate Client ID from user
	clientId := promptInput("Enter your Client Id: ", validate)

	// Get Client Secret from user (no format validation)
	clientSecret := promptInput("Enter your Client Secret: ", validate)

	configuration, _ := WriteSecretsToHomeDirectory(&Config{
		ClientId:     clientId,
		ClientSecret: clientSecret,
	})

	return configuration
}
