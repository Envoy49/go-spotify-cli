package prompt

import (
	"fmt"
	"github.com/manifoldco/promptui"
	"github.com/sirupsen/logrus"
	"go-spotify-cli/config"
	"go-spotify-cli/loader"
	"os"
	"regexp"
)

func EnvVarsPrompt() {
	if len(config.GlobalConfig.ClientSecret) > 0 || len(config.GlobalConfig.ClientId) > 0 {
		return
	}

	if config.VerifyConfigExists() {
		return
	}

	loader.Stop()

	SetupPrompt()

	validate := func(input string) error {
		// Regular expression for validating Spotify Client ID and Client Secret
		var clientIDRegex = regexp.MustCompile(`^[a-zA-Z0-9]{32}$`)

		if !clientIDRegex.MatchString(input) {
			return fmt.Errorf("invalid format")
		}
		return nil
	}

	promptClientId := promptui.Prompt{
		Label:    "Enter your Client Id",
		Validate: validate,
	}

	clientId, err := promptClientId.Run()
	if err != nil {
		logrus.WithError(err).Error("Client Id Prompt failed")
		os.Exit(1)
	}

	promptClientSecret := promptui.Prompt{
		Label: "Enter your Client Secret",
	}

	clientSecret, err := promptClientSecret.Run()
	if err != nil {
		logrus.WithError(err).Error("Client Secret Prompt failed")
		os.Exit(1)
	}

	config.WriteSecretsToHomeDirectory(clientSecret, clientId)
}
