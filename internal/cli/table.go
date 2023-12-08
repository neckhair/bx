package cli

import (
	"fmt"

	"github.com/charmbracelet/lipgloss"
	"github.com/charmbracelet/lipgloss/table"
)

func newTable() *table.Table {
	return table.New().
		Border(lipgloss.NormalBorder()).
		BorderStyle(lipgloss.NewStyle().Foreground(lipgloss.Color("99")))
}

func PrintTable(headers []string, rows [][]string) {
	t := newTable().Headers(headers...).Rows(rows...)
	fmt.Println(t)
}
