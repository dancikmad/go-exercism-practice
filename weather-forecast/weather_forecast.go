// Package weather to forecast thee current weather condition
// of various cities in Goblinicus.
package weather

var (
	// CurrentCondition represents the current weather forecast.
	CurrentCondition string
	// CurrentLocation represents the current city-location in Goblinicus.
	CurrentLocation string
)

// Forecast returns a string value representing the weather forecast in current location in Goblinicus.
func Forecast(city, condition string) string {
	CurrentLocation, CurrentCondition = city, condition
	return CurrentLocation + " - current weather condition: " + CurrentCondition
}
