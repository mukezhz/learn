package fanoutpattern

import (
	"fmt"
	"time"
)

type Producer struct {
	pubSub *Dispatcher
}

func NewProducer(pubSub *Dispatcher) *Producer {
	return &Producer{pubSub}
}

func (p *Producer) Produce() {
	for i := 0; ; i++ {
		p.pubSub.Broadcast(fmt.Sprintf("Message %d", i))
		time.Sleep(time.Second)
	}
}
