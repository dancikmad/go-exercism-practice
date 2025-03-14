package cards

// FavoriteCards returns a slice with the cards 2, 6 and 9 in that order.
func FavoriteCards() []int {
	cards := []int{2, 6, 9}
	return cards
}

// GetItem retrieves an item from a slice at given position.
// If the index is out of range, we want it to return -1.
func GetItem(slice []int, index int) int {
	return checkOutOfBound(slice, index)
}

// SetItem writes an item to a slice at given position overwriting an existing value.
// If the index is out of range the value needs to be appended.
func SetItem(slice []int, index, value int) []int {
	if index >= 0 && index < len(slice) {
		slice[index] = value // Update existing index
	} else {
		slice = append(slice, value) // Append to the slice if index is out of range
	}
	return slice
}

// PrependItems adds an arbitrary number of values at the front of a slice.
func PrependItems(slice []int, values ...int) []int {
	return append(values, slice...)
}

// RemoveItem removes an item from a slice by modifying the existing slice.
func RemoveItem(slice []int, index int) []int {
	if checkOutOfBound(slice, index) == -1 {
		return slice
	}
	return append(slice[:index], slice[index+1:]...)
}

func checkOutOfBound(slice []int, index int) int {
	if index >= len(slice) || index < 0 {
		return -1 // Return -1 when out of bounds
	}
	return slice[index]
}
