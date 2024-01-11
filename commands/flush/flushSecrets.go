package flush

import (
	"github.com/envoy49/go-spotify-cli/config"
	"github.com/spf13/cobra"
)

var FlushSecretsCommand = &cobra.Command{
	Use:   "flush-secrets",
	Short: "Flush Secrets",
	Run: func(cmd *cobra.Command, args []string) {
		config.WriteSecretsToHomeDirectory(nil) // fix this
	},
}
