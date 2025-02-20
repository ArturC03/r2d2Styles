package main

import "github.com/charmbracelet/lipgloss"

// Tag styles
var (
	errorStyle = lipgloss.NewStyle().
			Bold(true).
			Background(lipgloss.Color("#BF1A2F")).
			Foreground(lipgloss.Color("#ffffff")).
			Padding(0, 1)
	infoStyle = lipgloss.NewStyle().
			Bold(true).
			Background(lipgloss.Color("#454e9e")).
			Foreground(lipgloss.Color("#ffffff")).
			Padding(0, 1)
	warningStyle = lipgloss.NewStyle().
			Bold(true).
			Background(lipgloss.Color("#FFD700")).
			Foreground(lipgloss.Color("#ffffff")).
			Padding(0, 1)
)

// Colors
var (
	subtle     = lipgloss.AdaptiveColor{Light: "#D9DCCF", Dark: "#6E6E6E"}
	highlight  = lipgloss.AdaptiveColor{Light: "#874BFD", Dark: "#7D56F4"}
	special    = lipgloss.AdaptiveColor{Light: "#43BF6D", Dark: "#73F59F"}
	warning    = lipgloss.AdaptiveColor{Light: "#FFD700", Dark: "#FFC107"} // Amarelo para avisos
	errorColor = lipgloss.AdaptiveColor{Light: "#BF1A2F", Dark: "#D72638"} // Vermelho para erros
	info       = lipgloss.AdaptiveColor{Light: "#454e9e", Dark: "#3A57D6"} // Azul para informações
	accent     = lipgloss.AdaptiveColor{Light: "#FF5733", Dark: "#FF6B6B"} // Laranja para destaque
	background = lipgloss.AdaptiveColor{Light: "#FAFAFA", Dark: "#1E1E1E"} // Fundo adaptativo
)

// Help menu styles
var (
	baseStyle = lipgloss.NewStyle().
			BorderStyle(lipgloss.RoundedBorder()).
			BorderForeground(subtle).
			Padding(1, 2)
	titleStyle = lipgloss.NewStyle().
			Foreground(highlight).
			Bold(true).
			Padding(0, 1).
			MarginLeft(2)
	itemStyle = lipgloss.NewStyle().
			PaddingLeft(4).
			Foreground(lipgloss.Color("241"))
	selectedItemStyle = lipgloss.NewStyle().
				PaddingLeft(2).
				Foreground(special).
				Bold(true)
	messageStyle = lipgloss.NewStyle().
			Italic(true).
			Foreground(subtle).
			PaddingTop(1).
			PaddingLeft(2)
	infoItemStyle = lipgloss.NewStyle().
			PaddingLeft(2).
			Foreground(info).
			Bold(true)
)

func Tag(tagType string) string {
	switch tagType {
	case "error":
		return errorStyle.Render("ERROR")
	case "info":
		return infoStyle.Render("INFO")
	case "warning":
		return warningStyle.Render("WARNING")
	default:
		return ""
	}
}
