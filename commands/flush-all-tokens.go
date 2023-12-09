package commands

import (
	"github.com/spf13/cobra"
	"go-spotify-cli/config"
)

var FlushAllTokensCommand = &cobra.Command{
	Use:   "flush-all-tokens",
	Short: "Flush All Tokens",

	Run: func(cmd *cobra.Command, args []string) {
		config.WriteTokenToHomeDirectory(nil, false)
	},
}
