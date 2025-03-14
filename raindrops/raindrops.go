package raindrops

import "fmt"

// Convert converts a number into its corresponding raindrop sounds.
func Convert(number int) string {
	result := ""

	if number%3 == 0 {
		result += "Pling"
	}
	if number%5 == 0 {
		result += "Plang"
	}
	if number%7 == 0 {
		result += "Plong"
	}

	if result == "" {
		// If no raindrop sounds were added, return the number itself.
		return fmt.Sprintf("%d", number)
	}

	return result
}
