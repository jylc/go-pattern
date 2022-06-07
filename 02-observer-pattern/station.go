package observer

type WeatherData struct {
	observers   []IObserver
	temperature float64
	humidity    float64
	pressure    float64
}

func (wd *WeatherData) register(observer IObserver) {
	wd.observers = append(wd.observers, observer)
}

func (wd *WeatherData) remove(observer IObserver) {
	i := 0
	for i = 0; i < len(wd.observers); i++ {
		if wd.observers[i] == observer {
			break
		}
	}
	if i >= 0 {
		wd.observers = append(wd.observers[:i], wd.observers[i+1:]...)
	}
}

func (wd *WeatherData) notify() {
	for i := 0; i < len(wd.observers); i++ {
		wd.observers[i].update(wd.temperature, wd.humidity, wd.pressure)
	}
}

func (wd *WeatherData) measurementsChanged() {
	wd.notify()
}

func (wd *WeatherData) setMeasurements(temperature, humidity, pressure float64) {
	wd.temperature = temperature
	wd.humidity = humidity
	wd.pressure = pressure
	wd.measurementsChanged()
}
func NewWeatherData() *WeatherData {
	return &WeatherData{
		observers: make([]IObserver, 0),
	}
}
