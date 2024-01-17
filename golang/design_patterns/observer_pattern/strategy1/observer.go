package main

import "fmt"

type Observer interface {
	Update()
}

type ConcreteObserver struct {
	observable *ConcreteObservable
}

func NewConcreteObserver(observable *ConcreteObservable) *ConcreteObserver {
	return &ConcreteObserver{
		observable: observable,
	}
}

func (co *ConcreteObserver) Update() {
	fmt.Println("PUBLISHER STATE CHANGED:", co.observable.GetState())
}

type Display interface {
	Display()
}

func (co *ConcreteObserver) Display() {
	fmt.Println("STATE:", co.observable.GetState())
}
