package observer

type ISubject interface {
	register(observer IObserver)
	remove(observer IObserver)
	notify()
}
