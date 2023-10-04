package common

import "github.com/charmbracelet/lipgloss"

var HeaderStyle = lipgloss.NewStyle().
	Foreground(lipgloss.Color("#101030")).
	Background(lipgloss.Color("#51e2f5")).
	MarginBottom(1).
	MarginTop(1).
	Bold(true)

var ValueStyle = lipgloss.NewStyle().
	Foreground(lipgloss.Color("#51e2f5")).
	Bold(true)

var BoxStyle = lipgloss.NewStyle()
