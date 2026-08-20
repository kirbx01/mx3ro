package ui

import (
	"strings"
	"github.com/fatih/color"
)

func LineColor(line string) func(format string, a ...interface{}) {
	switch {
	case strings.Contains(line, "Red"):
		return color.New(color.FgRed).PrintfFunc()
	case strings.Contains(line, "Yellow"):
		return color.New(color.FgYellow).PrintfFunc()
	case strings.Contains(line, "Blue"):
		return color.New(color.FgBlue).PrintfFunc()
	case strings.Contains(line, "Green"):
		return color.New(color.FgGreen).PrintfFunc()
	case strings.Contains(line, "Violet"):
		return color.New(color.FgMagenta).PrintfFunc()
	case strings.Contains(line, "Pink"):
		return color.New(color.FgHiMagenta).PrintfFunc()
	case strings.Contains(line, "Magenta"):
		return color.New(color.FgHiRed).PrintfFunc()
	case strings.Contains(line, "Orange"):
		return color.New(color.FgHiYellow).PrintfFunc()
	case strings.Contains(line, "Gray"), strings.Contains(line, "Rapid"), strings.Contains(line, "Aqua"):
		return color.New(color.FgCyan).PrintfFunc()
	default:
		return color.New(color.FgWhite).PrintfFunc()
	}
}