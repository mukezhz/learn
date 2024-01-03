package fanoutpattern

import "fmt"

type Subscriber struct {
	name string
	msg  chan string
}

func NewSubscriber(name string) *Subscriber {
	return &Subscriber{
		name: name,
		msg:  make(chan string),
	}
}

func (s *Subscriber) Listen() {
	for msg := range s.msg {
		fmt.Printf("%v -> received: %v\n", s.name, msg)
	}
}
