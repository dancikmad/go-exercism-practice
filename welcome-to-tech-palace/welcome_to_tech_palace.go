package techpalace

import (
	"fmt"
	"strings"
)

// WelcomeMessage returns a welcome message for the customer.
func WelcomeMessage(customer string) string {
	return "Welcome to the Tech Palace, " + strings.ToUpper(customer)
}

// AddBorder adds a border to a welcome message.
func AddBorder(welcomeMsg string, numStarsPerLine int) string {
	starLines := strings.Repeat("*", numStarsPerLine)
	return fmt.Sprintf("%s\n%s\n%s", starLines, welcomeMsg, starLines)
}

// CleanupMessage cleans up an old marketing message.
func CleanupMessage(oldMsg string) string {
	cleanupMsg := strings.ReplaceAll(oldMsg, "*", "")

	// Trim whitespaces and newlines
	cleanupMsg = strings.TrimSpace(cleanupMsg)

	return cleanupMsg
}
