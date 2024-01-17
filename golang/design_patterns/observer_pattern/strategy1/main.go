package main

func main() {
	observable := NewConcreteObservable()
	observer := NewConcreteObserver(observable)
	observable.AddObserver(observer)
	observable.SetState(1)
	observable.NotifyObservers()
	observable.SetState(2)
	observable.NotifyObservers()
}
