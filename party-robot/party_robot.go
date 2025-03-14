package partyrobot

import "fmt"

// Welcome greets a person by name.
func Welcome(name string) string {
	return fmt.Sprintf("Welcome to my party, %s!", name)
}

// HappyBirthday wishes happy birthday to the birthday person and exclaims their age.
func HappyBirthday(name string, age int) string {
	return fmt.Sprintf("Happy birthday %s! You are now %d years old!", name, age)
}

// AssignTable assigns a table to each guest.
func AssignTable(name string, table int, neighbor, direction string, distance float64) string {
	// Format the table number so it can be 3 digits with leading zeros
	formattedTable := fmt.Sprintf("%03d", table)

	// Format the distance to one decimal place
	formattedDistance := fmt.Sprintf("%.1f", distance)

	// Construct the full message
	message := fmt.Sprintf(
		"Welcome to my party, %s!\nYou have been assigned to table %s. Your table is %s, exactly %s meters from here.\nYou will be sitting next to %s.",
		name,
		formattedTable,
		direction,
		formattedDistance,
		neighbor,
	)
	return message
}
