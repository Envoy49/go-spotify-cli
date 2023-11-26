package config

import (
	"bufio"
	"fmt"
	"go-spotify-cli/spinnerInstance"
	"os"
	"strings"
)

func EnvVarsPrompt() {
	if len(GlobalConfig.ClientSecret) > 0 || len(GlobalConfig.ClientId) > 0 {
		return
	}

	if VerifyConfigExists() {
		return
	}

	spinnerInstance.Stop()

	PrintPromt()

	reader := bufio.NewReader(os.Stdin)

	fmt.Print("Enter client id: ")

	clientId, _ := reader.ReadString('\n')

	clientId = strings.TrimSpace(clientId)

	fmt.Print("Enter client secret: ")

	clientSecret, _ := reader.ReadString('\n')

	clientSecret = strings.TrimSpace(clientSecret)

	WriteToHomeDirectory(clientSecret, clientId)
}
