package main

import (
	"fmt"

	"github.com/charmbracelet/lipgloss"
)

var style = lipgloss.NewStyle().
	Bold(true).
	Foreground(lipgloss.Color("#FAFAFA")).
	Background(lipgloss.AdaptiveColor{Light: "#444444", Dark: "#888888"}).
	PaddingTop(2).
	PaddingLeft(4).
	Width(22)

func main() {
	fmt.Println(style.Render("Hello, kitty"))
}
