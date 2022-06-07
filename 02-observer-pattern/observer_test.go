package observer

import "testing"

func TestObserverPattern(t *testing.T) {
	weatherData := NewWeatherData()
	NewCurrentConditionDisplay(weatherData)
	weatherData.setMeasurements(80, 65, 30.4)
	weatherData.setMeasurements(82, 70, 29.2)
}
