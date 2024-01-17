package main

import "fmt"

type Observable interface {
	AddObserver(o Observer)
	RemoveObserver(o Observer)
	NotifyObservers()
}

type ConcreteObservable struct {
	observers []Observer
	state     int
}

func NewConcreteObservable() *ConcreteObservable {
	return &ConcreteObservable{
		observers: make([]Observer, 0),
		state:     0,
	}
}

func (co *ConcreteObservable) AddObserver(o Observer) {
	co.observers = append(co.observers, o)
}

func (co *ConcreteObservable) RemoveObserver(o Observer) {
	for i, observer := range co.observers {
		if observer == o {
			co.observers = append(co.observers[:i], co.observers[i+1:]...)
		}
	}
}

func (co *ConcreteObservable) NotifyObservers() {
	for _, observer := range co.observers {
		fmt.Println("Notifying state changed:", co.state)
		observer.Update()
	}
}

func (co *ConcreteObservable) SetState(state int) {
	co.state = state
}

func (co *ConcreteObservable) GetState() int {
	return co.state
}
