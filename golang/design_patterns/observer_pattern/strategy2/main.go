package main

func main() {
	observable := NewConcreteObservable()
	observer := NewConcreteObserver()
	observable.AddObserver(observer)
	observable.SetState(1)
	observable.NotifyObservers()
	observer.Display()
	observable.SetState(2)
	observable.NotifyObservers()
	observer.Display()
}
