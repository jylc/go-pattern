package observer

import "fmt"

type CurrentConditionDisplay struct {
	weatherData ISubject
	temperature float64
	humidity    float64
}

func (d *CurrentConditionDisplay) display() {
	out := fmt.Sprintf("Current condition: %fF degrees and %f humidity", d.temperature, d.humidity)
	fmt.Println(out)
}

func (d *CurrentConditionDisplay) update(temperature, humidity, pressure float64) {
	d.temperature = temperature
	d.humidity = humidity
	d.display()
}

func NewCurrentConditionDisplay(weatherData ISubject) *CurrentConditionDisplay {
	d := &CurrentConditionDisplay{
		weatherData: weatherData,
	}
	weatherData.register(d)
	return d
}
