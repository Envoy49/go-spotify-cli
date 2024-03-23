package config

import (
	"fmt"
	"os"
	"os/exec"
	"runtime"

	"github.com/charmbracelet/lipgloss"
	"golang.org/x/term"
)

func SecretsSetupPrompt() {
	clearTerminal()
	// Get terminal width 
	tw, _, _ := term.GetSize(int(os.Stdout.Fd()))

	// Header box style
	headerBoxStyle := lipgloss.NewStyle().
		Foreground(lipgloss.Color("black")).
		Background(lipgloss.Color("#51e2f5")).
		Bold(true).
		PaddingTop(1).
		PaddingBottom(1).
		PaddingLeft(3).
		PaddingRight(3).
		Align(lipgloss.Center).
		Border(lipgloss.RoundedBorder()).
		BorderForeground(lipgloss.Color("#51e2f5"))

	// Apply a larger font size effect
	headerContent := "GO Spotify CLI Setup Guide"
	header := headerBoxStyle.Render(headerContent)

	// Body style
	bodyStyle := lipgloss.NewStyle().
		Bold(true).
		Foreground(lipgloss.Color("#51e2f5")).
		PaddingLeft(4).
		PaddingRight(4).
		PaddingTop(1).
		PaddingBottom(1)

	// URL style
	urlStyle := lipgloss.NewStyle().
		Bold(true).
		Foreground(lipgloss.Color("#007AFF")) // Blue color for URLs

	// URLs
	url1 := urlStyle.Render(fmt.Sprintf("%s%s", ServerUrl, UserModifyPlaybackStateRouteCallback))
	url2 := urlStyle.Render(fmt.Sprintf("%s%s", ServerUrl, UserReadPlaybackStateRouteCallback))
	url3 := urlStyle.Render(fmt.Sprintf("%s%s", ServerUrl, UserLibraryReadRouteCallback))

	// Prepare body and footer content with highlighted URLs
	body := bodyStyle.Render(fmt.Sprintf(`To get started, you'll need a 'Client ID' and 'Client Secret' from Spotify's Developer Dashboard:

1. 🔗 Navigate to: https://developer.spotify.com/dashboard/applications

2. 🚪 Sign in or create a Spotify account.

3. ➕ Click on 'Create An App'.

4. 📜 Fill in the app details.

5. ❗ In the app settings, set your Redirect URIs. Ensure your CLI tool's callback URL is added.

6. 🌐 In order to authenticate with Spotify, in Redirect URIs field please enter following URLs in the app you created:

   📎 %s 
      %s 
      %s

7. 🛍 Once App created, you'll find the 'Client ID' and 'Client Secret' on the app details page.`, url1, url2, url3))

	// Footer style
	footerStyle := lipgloss.NewStyle().
		Bold(true).
		Foreground(lipgloss.Color("#FF0000")).
		PaddingTop(1).
		PaddingBottom(1)

	footer := footerStyle.Render("🚫 Remember: Keep your 'Client Secret and Client Id' confidential. Never share it!")

	// Determine the widest element for the overall box
	maxWidth := lipgloss.Width(header)
	bodyWidth := lipgloss.Width(body)
	if bodyWidth > maxWidth {
		maxWidth = bodyWidth
	}
	footerWidth := lipgloss.Width(footer)
	if footerWidth > maxWidth {
		maxWidth = footerWidth
	}

	// Box style for the entire content
	box := lipgloss.NewStyle().
		Border(lipgloss.RoundedBorder()).
		BorderForeground(lipgloss.Color("#51e2f5")).
		Width(tw-2).
		MarginTop(1).
		MarginBottom(1)

	// Render the box with centered content
	fmt.Println(box.Render(lipgloss.JoinVertical(lipgloss.Center, header, body, footer)))
}

func clearTerminal() {
	var cmd *exec.Cmd
	if runtime.GOOS == "windows" {
		cmd = exec.Command("cmd", "/c", "cls")
	} else {
		cmd = exec.Command("clear")
	}

	cmd.Stdout = os.Stdout
	cmd.Run()
}
