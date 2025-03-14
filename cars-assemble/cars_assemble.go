package cars

// CalculateWorkingCarsPerHour calculates how many working cars are
// produced by the assembly line every hour.
func CalculateWorkingCarsPerHour(productionRate int, successRate float64) float64 {
	// panic("CalculateWorkingCarsPerHour not implemented")
	return float64(productionRate) * (successRate / 100)
}

// CalculateWorkingCarsPerMinute calculates how many working cars are
// produced by the assembly line every minute.
func CalculateWorkingCarsPerMinute(productionRate int, successRate float64) int {
	successHourRate := CalculateWorkingCarsPerHour(productionRate, successRate)
	return int(successHourRate) / 60
}

// CalculateCost works out the cost of producing the given number of cars.
func CalculateCost(carsCount int) uint {
	bulkGroups := carsCount / 10
	individualCars := carsCount % 10

	totalCost := (bulkGroups * 95000) + (individualCars * 10000)
	return uint(totalCost)
}
