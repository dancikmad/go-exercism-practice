package birdwatcher

import "fmt"

// TotalBirdCount return the total bird count by summing
// the individual day's counts.
func TotalBirdCount(birdsPerDay []int) int {
	var totalBirds int
	for i := 0; i < len(birdsPerDay); i++ {
		totalBirds += birdsPerDay[i]
	}
	return totalBirds
}

// BirdsInWeek returns the total bird count by summing
// only the items belonging to the given week.
func BirdsInWeek(birdsPerDay []int, week int) int {
	daysPerWeek := 7
	startIndex := (week - 1) * daysPerWeek
	endIndex := startIndex + daysPerWeek

	if startIndex >= len(birdsPerDay) {
		fmt.Println("Week", week, "is out of range")
		return 0
	}
	if endIndex > len(birdsPerDay) {
		endIndex = len(birdsPerDay)
	}

	total := 0
	for i := startIndex; i < endIndex; i++ {
		total += birdsPerDay[i]
	}
	return total
}

// FixBirdCountLog returns the bird counts after correcting
// the bird counts for alternate days.
func FixBirdCountLog(birdsPerDay []int) []int {
	for i := 0; i < len(birdsPerDay); i++ {
		if i%2 == 0 {
			birdsPerDay[i] += 1
		}
	}
	return birdsPerDay
}
