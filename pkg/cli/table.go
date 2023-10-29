package cli

import (
	"fmt"
	"os"

	"github.com/charmbracelet/lipgloss"
	"github.com/charmbracelet/lipgloss/table"
)

var (
	re          = lipgloss.NewRenderer(os.Stdout)
	baseStyle   = re.NewStyle().Padding(0, 1)
	headerStyle = baseStyle.Copy().Foreground(lipgloss.Color("252")).Bold(true)
)

func printTable(headers []string, values [][]string) {
	t := table.New().
		Border(lipgloss.NormalBorder()).
		BorderStyle(re.NewStyle().Foreground(lipgloss.Color("238"))).
		Headers(headers...).
		Width(100).
		Rows(values...).
		StyleFunc(func(row, col int) lipgloss.Style {
			if row == 0 {
				return headerStyle
			}

			if row%2 == 0 {
				return baseStyle.Copy().Foreground(lipgloss.Color("245"))
			}
			return baseStyle.Copy().Foreground(lipgloss.Color("252"))
		})
	fmt.Println(t)
}
