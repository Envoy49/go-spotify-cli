package commands

import (
	"github.com/spf13/cobra"
	"go-spotify-cli/config"
)

var FlushTokensCommand = &cobra.Command{
	Use:   "flush-tokens",
	Short: "Flush Tokens",
	Run: func(cmd *cobra.Command, args []string) {
		config.WriteTokenToHomeDirectory(nil, false)
	},
}
