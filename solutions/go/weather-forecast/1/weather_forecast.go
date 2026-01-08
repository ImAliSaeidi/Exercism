// Package weather implements a weather forecast service.
package weather

var (
	//CurrentCondition represent current weather condition.
	CurrentCondition string

	//CurrentLocation represent current location to forecast.
	CurrentLocation string
)

// Forecast returns a string representing the weather forecast for the given city and condition.
func Forecast(city, condition string) string {
	CurrentLocation, CurrentCondition = city, condition
	return CurrentLocation + " - current weather condition: " + CurrentCondition
}
