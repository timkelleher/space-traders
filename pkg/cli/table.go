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
		BorderStyle(lipgloss.NewStyle().Foreground(lipgloss.Color("99"))).
		Headers(headers...).
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

func printError(err error) {
	printTable([]string{"Error Message"}, [][]string{{err.Error()}})
}
