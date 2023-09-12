package config

import (
	"bufio"
	"fmt"
	"github.com/spf13/cobra"
	"os"
	"strings"
)

func EnvVarsPrompt(cmd *cobra.Command, args []string) {
	if len(GlobalConfig.ClientSecret) > 0 || len(GlobalConfig.ClientId) > 0 {
		return
	}

	if VerifyConfigExists() {
		return
	}

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
