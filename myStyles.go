package r2d2Styles

import "github.com/charmbracelet/lipgloss"

// Colors - Todas as cores agora são adaptáveis
var (
	subtle     = lipgloss.AdaptiveColor{Light: "#B0B7C3", Dark: "#6E6E6E"} // More subtle, neutral
	highlight  = lipgloss.AdaptiveColor{Light: "#7F66FF", Dark: "#9D7EFF"} // Lighter purple, great contrast
	special    = lipgloss.AdaptiveColor{Light: "#3DCA5E", Dark: "#53D57C"} // Green with pop
	warning    = lipgloss.AdaptiveColor{Light: "#FFB92D", Dark: "#FFAE3D"} // Brighter yellow
	errorColor = lipgloss.AdaptiveColor{Light: "#FF4F5A", Dark: "#D93242"} // Brighter red
	info       = lipgloss.AdaptiveColor{Light: "#4C6EF5", Dark: "#3A5BDF"} // Clean blue
	accent     = lipgloss.AdaptiveColor{Light: "#FF5733", Dark: "#FF6B6B"} // Vibrant orange
	background = lipgloss.AdaptiveColor{Light: "#F0F4F8", Dark: "#2E2E2E"} // Soft light and dark background
	success    = lipgloss.AdaptiveColor{Light: "#28A745", Dark: "#218838"} // Verde adaptável

	// Cores adicionais para text readability
	primaryText   = lipgloss.AdaptiveColor{Light: "#000000", Dark: "#FFFFFF"} // Texto principal
	secondaryText = lipgloss.AdaptiveColor{Light: "#666666", Dark: "#CCCCCC"} // Texto secundário
	invertedText  = lipgloss.AdaptiveColor{Light: "#FFFFFF", Dark: "#000000"} // Texto invertido para backgrounds coloridos
	selectedBg    = lipgloss.AdaptiveColor{Light: "#E8E8E8", Dark: "#3A3A3A"} // Background para itens selecionados
	itemText      = lipgloss.AdaptiveColor{Light: "#555555", Dark: "#AAAAAA"} // Texto para itens normais
)

// Tag styles - Agora totalmente responsivos
var (
	boldStyle = lipgloss.NewStyle().
			Bold(true).
			Foreground(highlight)

	errorStyle = lipgloss.NewStyle().
			Bold(true).
			Background(errorColor).
			Foreground(primaryText).
			Padding(0, 1)

	infoStyle = lipgloss.NewStyle().
			Bold(true).
			Background(info).
			Foreground(primaryText).
			Padding(0, 1)

	warningStyle = lipgloss.NewStyle().
			Bold(true).
			Background(warning).
			Foreground(primaryText).
			Padding(0, 1)

	successStyle = lipgloss.NewStyle().
			Bold(true).
			Background(success).
			Foreground(primaryText).
			Padding(0, 1)
)

// Help menu styles - Totalmente responsivos
var (
	baseStyle = lipgloss.NewStyle().
			BorderStyle(lipgloss.RoundedBorder()).
			BorderForeground(subtle).
			Padding(1, 3)

	titleStyle = lipgloss.NewStyle().
			Foreground(highlight).
			Bold(true).
			Padding(0, 2).
			MarginLeft(3).
			Italic(true)

	itemStyle = lipgloss.NewStyle().
			PaddingLeft(4).
			Foreground(itemText).
			Italic(true)

	selectedItemStyle = lipgloss.NewStyle().
				PaddingLeft(2).
				Foreground(special).
				Bold(true).
				Background(selectedBg)

	messageStyle = lipgloss.NewStyle().
			Italic(true).
			Foreground(subtle).
			PaddingTop(1).
			PaddingLeft(3).
			MarginTop(1)

	infoItemStyle = lipgloss.NewStyle().
			PaddingLeft(3).
			Foreground(info).
			Bold(true)
)

// Função Tag mantém-se igual
func Tag(tagType string) string {
	switch tagType {
	case "error":
		return errorStyle.Render("ERROR")
	case "info":
		return infoStyle.Render("INFO")
	case "warning":
		return warningStyle.Render("WARNING")
	case "success":
		return successStyle.Render("SUCCESS")
	default:
		return ""
	}
}

// Funções auxiliares para acesso direto aos estilos
func GetBoldStyle() lipgloss.Style {
	return boldStyle
}

func GetErrorStyle() lipgloss.Style {
	return errorStyle
}

func GetInfoStyle() lipgloss.Style {
	return infoStyle
}

func GetWarningStyle() lipgloss.Style {
	return warningStyle
}

func GetSuccessStyle() lipgloss.Style {
	return successStyle
}

func GetBaseStyle() lipgloss.Style {
	return baseStyle
}

func GetTitleStyle() lipgloss.Style {
	return titleStyle
}

func GetItemStyle() lipgloss.Style {
	return itemStyle
}

func GetSelectedItemStyle() lipgloss.Style {
	return selectedItemStyle
}

func GetMessageStyle() lipgloss.Style {
	return messageStyle
}

func GetInfoItemStyle() lipgloss.Style {
	return infoItemStyle
}
