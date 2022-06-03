package observer

type IObserver interface {
	update(temp, humidity, pressure float64)
}
