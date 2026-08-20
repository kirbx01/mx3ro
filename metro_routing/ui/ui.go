package ui

import (
	"strings"

	"github.com/charmbracelet/lipgloss"
)

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

func normalizeLineName(line string) string {
	trimmed := strings.TrimSpace(line)
	trimmed = strings.TrimSuffix(trimmed, " line")
	trimmed = strings.TrimSuffix(trimmed, " Line")
	trimmed = strings.TrimSuffix(trimmed, "LINE")
	return trimmed
}

func colorForLine(line string) lipgloss.Style {
	style := lipgloss.NewStyle()
	name := normalizeLineName(line)
	if colorHex, ok := lineColors[name]; ok {
		style = style.Foreground(lipgloss.Color(colorHex))
	}
	return style
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