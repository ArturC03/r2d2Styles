package main

import (
	"fmt"
	"github.com/charmbracelet/lipgloss"
)

// Print tag and message together
func PrintTagMessage(tagType, message string) {
	switch tagType {
	case "error":
		fmt.Println(errorStyle.Render("ERROR") + " " + message)
	case "info":
		fmt.Println(infoStyle.Render("INFO") + " " + message)
	case "warning":
		fmt.Println(warningStyle.Render("WARNING") + " " + message)
	default:
		fmt.Println(message) // Default plain text
	}
}

// Print semantic error with line & column info
func PrintSemanticError(line, column int, message string) {
	errorMsg := fmt.Sprintf("Line %d, Column %d: %s", line, column, message)
	fmt.Println(errorStyle.Render("SEMANTIC ERROR") + " " + errorMsg)
}

// Print a message during code generation
func PrintCodeGenInfo(phase, message string) {
	phaseStyle := lipgloss.NewStyle().
		Bold(true).
		Background(lipgloss.Color("#008080")). // Teal for codegen phases
		Foreground(lipgloss.Color("#ffffff")).
		Padding(0, 1)

	fmt.Println(phaseStyle.Render(phase) + " " + message)
}

// Print success message
func PrintSuccess(message string) {
	successStyle := lipgloss.NewStyle().
		Bold(true).
		Background(lipgloss.Color("#43BF6D")). // Green success
		Foreground(lipgloss.Color("#ffffff")).
		Padding(0, 1)

	fmt.Println(successStyle.Render("SUCCESS") + " " + message)
}

// Print debug messages
func PrintDebug(message string) {
	debugStyle := lipgloss.NewStyle().
		Italic(true).
		Foreground(lipgloss.Color("#6E6E6E")) // Grey debug messages

	fmt.Println(debugStyle.Render("[DEBUG] " + message))
}

// Example usage
func main() {
	PrintTagMessage("info", "Starting semantic analysis...")
	PrintSemanticError(12, 5, "Undefined variable 'x'")
	PrintCodeGenInfo("Parsing", "Analyzing AST structure")
	PrintCodeGenInfo("Generating", "Creating intermediate representation")
	PrintSuccess("Code generation completed!")
	PrintDebug("Checking variable scope for optimizations")
}
