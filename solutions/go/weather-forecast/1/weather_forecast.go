//Package weather tlls you about the weather.
package weather

var (
    //CurrentCondition ...
	CurrentCondition string
    //CurrentLocation ....
	CurrentLocation  string
)

// Forecast is something.
func Forecast(city, condition string) string {
	CurrentLocation, CurrentCondition = city, condition
	return CurrentLocation + " - current weather condition: " + CurrentCondition
}
