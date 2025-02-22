package r2d2Styles

import "fmt"

// Function to generate an info message
func InfoMessage(message string) string {
	return fmt.Sprintf("%s %s", Tag("info"), message)
}

// Function to generate a warning message
func WarningMessage(message string) string {
	return fmt.Sprintf("%s %s", Tag("warning"), message)
}

// Function to generate an error message
func ErrorMessage(message string) string {
	return fmt.Sprintf("%s %s", Tag("error"), message)
}
