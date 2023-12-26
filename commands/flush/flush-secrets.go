package flush

import (
	"github.com/spf13/cobra"
	"go-spotify-cli/config"
)

var FlushSecretsCommand = &cobra.Command{
	Use:   "flush-secrets",
	Short: "Flush Secrets",
	Run: func(cmd *cobra.Command, args []string) {
		config.WriteSecretsToHomeDirectory("", "")
	},
}
