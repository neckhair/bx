package cli

import (
	"fmt"

	"github.com/charmbracelet/lipgloss"
)

func PrintError(msg string) {
	style := lipgloss.NewStyle().Foreground(lipgloss.Color("#990000"))
	fmt.Println(style.Render(msg))
}

func Bold(msg string) {
	style := lipgloss.NewStyle().Bold(true)
	fmt.Println(style.Render(msg))
}
