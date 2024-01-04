package flush

import (
	"github.com/envoy49/go-spotify-cli/config"
	"github.com/spf13/cobra"
)

var FlushTokensCommand = &cobra.Command{
	Use:   "flush-tokens",
	Short: "Flush Tokens",
	Run: func(cmd *cobra.Command, args []string) {
		config.WriteTokenToHomeDirectory(nil, false)
	},
}
