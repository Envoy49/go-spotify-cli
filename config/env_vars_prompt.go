package config

import (
	"fmt"
	"github.com/manifoldco/promptui"
	"github.com/sirupsen/logrus"
	"go-spotify-cli/spinnerInstance"
	"regexp"
)

func EnvVarsPrompt() {
	if len(GlobalConfig.ClientSecret) > 0 || len(GlobalConfig.ClientId) > 0 {
		return
	}

	if VerifyConfigExists() {
		return
	}

	spinnerInstance.Stop()

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
		return
	}

	promptClientSecret := promptui.Prompt{
		Label: "Enter your Client Secret",
	}

	clientSecret, err := promptClientSecret.Run()
	if err != nil {
		logrus.WithError(err).Error("Client Secret Prompt failed")
		return
	}

	WriteToHomeDirectory(clientSecret, clientId)
}
