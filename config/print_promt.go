package config

import (
	"fmt"
	"go-spotify-cli/constants"
)

func PrintPromt() {
	fmt.Println("==============================================")
	fmt.Println("          GO Spotify CLI Setup Guide          ")
	fmt.Println("==============================================")
	fmt.Println()
	fmt.Println("To get started, you'll need a `Client ID` and `Client Secret` from Spotify's Developer Dashboard:")
	fmt.Println()
	fmt.Println("1. 🔗 Navigate to: https://developer.spotify.com/dashboard/applications")
	fmt.Println()
	fmt.Println("2. 🚪 Sign in or create a Spotify account.")
	fmt.Println()
	fmt.Println("3. ➕ Click on 'Create An App'.")
	fmt.Println()
	fmt.Println("4. 📜 Fill in the app details.")
	fmt.Println()
	fmt.Println("5. ❗ In the app settings, set your Redirect URIs. Ensure your CLI tool's callback URL is added.")
	fmt.Println()
	fmt.Println("6. 🌐 In order to authenticate with Spotify, please enter following URLs in the app you created:")
	fmt.Println()
	fmt.Printf("   📎 %s%s and %s%s\n",
		constants.ServerUrl, constants.UserModifyPlaybackStateRouteCallback,
		constants.ServerUrl, constants.UserReadPlaybackStateRouteCallback)
	fmt.Println()
	fmt.Println("7. 🛍  Once App created, you'll find the `Client ID` and `Client Secret` on the app details page.")
	fmt.Println()
	fmt.Println("==============================================")
	fmt.Println("🚫 Remember: Keep your `Client Secret and Client Id` confidential. Never share it!")
	fmt.Println("==============================================")
}
