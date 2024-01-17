package main

import "fmt"

type Observer interface {
	Update(o Observable)
}

type ConcreteObserver struct {
	State int
}

func NewConcreteObserver() *ConcreteObserver {
	return &ConcreteObserver{}
}

func (co *ConcreteObserver) Update(o Observable) {
	fmt.Println("PUBLISHER STATE CHANGED:", o.GetState())
	co.State = o.GetState()
}

type Display interface {
	Display()
}

func (co *ConcreteObserver) Display() {
	fmt.Println("STATE:", co.State)
}
