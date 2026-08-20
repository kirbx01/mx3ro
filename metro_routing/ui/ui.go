package ui

import (
	"strings"

	"github.com/charmbracelet/lipgloss"
	"github.com/muesli/termenv"
)

func init() {
	lipgloss.SetColorProfile(termenv.TrueColor)
}

var lineColors = map[string]string{
	"Red":     "#E24B4A",
	"Yellow":  "#EF9F27",
	"Blue":    "#378ADD",
	"Green":   "#639922",
	"Violet":  "#7F77DD",
	"Pink":    "#D4537E",
	"Magenta": "#993556",
	"Orange":  "#D85A30",
	"Gray":    "#5F5E5A",
	"Aqua":    "#1D9E75",
	"Rapid":   "#1D9E75",
}

func colorForLine(line string) lipgloss.Style {
	if strings.HasPrefix(line, "interchange:") {
		parts := strings.Split(line, "->")
		if len(parts) == 2 {
			line = parts[1]
		}
	}

	for key, hex := range lineColors {
		if strings.Contains(line, key) {
			return lipgloss.NewStyle().Foreground(lipgloss.Color(hex))
		}
	}
	return lipgloss.NewStyle().Foreground(lipgloss.Color("#888880"))
}

func RenderRoute(stations []string, lines []string) string {
	var b strings.Builder

	for i, station := range stations {
		isSource := i == 0
		isDestination := i == len(stations)-1
		isInterchange := i > 0 && i < len(stations)-1 && lines[i] != lines[i-1]
		isKeyStation := isSource || isDestination || isInterchange

		style := colorForLine(lines[i])

		glyph := "○"
		if isKeyStation {
			glyph = "●"
		}

		textStyle := style
		if isKeyStation {
			textStyle = style.Bold(true)
		}

		b.WriteString(style.Render(glyph))
		b.WriteString(" ")
		b.WriteString(textStyle.Render(station))

		if i < len(stations)-1 {
			b.WriteString("  ")
			b.WriteString(style.Render("───"))
			b.WriteString("  ")
		}
	}

	return b.String()
}
