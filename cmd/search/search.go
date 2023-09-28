package search

//
//import (
//	"github.com/sirupsen/logrus"
//	"github.com/spf13/cobra"
//	"go-spotify-cli/cmd"
//	"go-spotify-cli/common"
//	"go-spotify-cli/constants"
//	"go-spotify-cli/server"
//)
//
//func search(accessToken string) {
//	params := &cmd.PlayerParams{
//		AccessToken: accessToken,
//		Method:      "GET",
//		Endpoint:    constants.SpotifySearchEndpoint,
//	}
//	_, err := cmd.FetchCommand(params)
//
//	if err != nil {
//		switch e := err.(type) {
//		case common.SpotifyAPIError:
//			if e.Detail.Error.Message == "Player command failed: No active device found" {
//				Device()
//			}
//
//		}
//
//		logrus.WithError(err).Error("Error pausing your track")
//
//	} else {
//		logrus.Info("Paused")
//	}
//}
//
//var PauseCommand = &cobra.Command{
//	Use:   "search",
//	Short: "Search spotify song",
//	Run: func(cmd *cobra.Command, args []string) {
//		token := server.ReadUserModifyTokenOrFetchFromServer()
//		search(token)
//	},
//}
