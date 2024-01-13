package flush

import (
	"github.com/envoy49/go-spotify-cli/config"
	"github.com/sirupsen/logrus"
	"github.com/spf13/cobra"
	"os"
)

func FlushSecretsCommand(fetchType *config.FetchType) *cobra.Command {
	return &cobra.Command{
		Use:   "flush-secrets",
		Short: "Flush Secrets",
		Run: func(cmd *cobra.Command, args []string) {
			if fetchType.NewFetch == true {
				os.Exit(0)
			} else {
				_, err := config.WriteSecretsToHomeDirectory(nil)
				if err != nil {
					logrus.WithError(err)
				}
			}
		},
	}
}
