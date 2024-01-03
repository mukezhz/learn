package fanoutpattern

import (
	"sync"
)

type Dispatcher struct {
	subscribers []*Subscriber
	mu          sync.Mutex
}

func NewDispatcher() *Dispatcher {
	return &Dispatcher{}
}

func (ps *Dispatcher) Subscribe(subscriber *Subscriber) {
	ps.mu.Lock()
	defer ps.mu.Unlock()

	ps.subscribers = append(ps.subscribers, subscriber)
	go subscriber.Listen()
}

func (ps *Dispatcher) Unsubscribe(subscriber *Subscriber) {
	ps.mu.Lock()
	defer ps.mu.Unlock()

	for i, s := range ps.subscribers {
		if s == subscriber {
			ps.subscribers = append(ps.subscribers[:i], ps.subscribers[i+1:]...)
			return
		}
	}
}

func (ps *Dispatcher) Broadcast(msg string) {
	ps.mu.Lock()
	defer ps.mu.Unlock()

	for _, subscriber := range ps.subscribers {
		subscriber.msg <- msg
	}
}

func (ps *Dispatcher) Close() {
	ps.mu.Lock()
	defer ps.mu.Unlock()

	for _, subscriber := range ps.subscribers {
		close(subscriber.msg)
	}
}
