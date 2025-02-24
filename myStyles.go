package r2d2Styles

import "github.com/charmbracelet/lipgloss"

// Tag styles
var (
	errorStyle = lipgloss.NewStyle().
			Bold(true).
			Background(lipgloss.Color("#FF4F5A")).
			Foreground(lipgloss.Color("#FFFFFF")).
			Padding(0, 1)

	infoStyle = lipgloss.NewStyle().
			Bold(true).
			Background(lipgloss.Color("#4C6EF5")).
			Foreground(lipgloss.Color("#FFFFFF")).
			Padding(0, 1)

	warningStyle = lipgloss.NewStyle().
			Bold(true).
			Background(lipgloss.Color("#FFB92D")).
			Foreground(lipgloss.Color("#000000")).
			Padding(0, 1)
)

// Colors
var (
	subtle     = lipgloss.AdaptiveColor{Light: "#B0B7C3", Dark: "#6E6E6E"} // More subtle, neutral
	highlight  = lipgloss.AdaptiveColor{Light: "#7F66FF", Dark: "#9D7EFF"} // Lighter purple, great contrast
	special    = lipgloss.AdaptiveColor{Light: "#3DCA5E", Dark: "#53D57C"} // Green with pop
	warning    = lipgloss.AdaptiveColor{Light: "#FFB92D", Dark: "#FFAE3D"} // Brighter yellow
	errorColor = lipgloss.AdaptiveColor{Light: "#FF4F5A", Dark: "#D93242"} // Brighter red
	info       = lipgloss.AdaptiveColor{Light: "#4C6EF5", Dark: "#3A5BDF"} // Clean blue
	accent     = lipgloss.AdaptiveColor{Light: "#FF5733", Dark: "#FF6B6B"} // Vibrant orange
	background = lipgloss.AdaptiveColor{Light: "#F0F4F8", Dark: "#2E2E2E"} // Soft light and dark background
)

// Help menu styles
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
			Foreground(lipgloss.Color("241")).
			Italic(true)

	selectedItemStyle = lipgloss.NewStyle().
				PaddingLeft(2).
				Foreground(special).
				Bold(true).
				Background(lipgloss.Color("235"))

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
