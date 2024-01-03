package main

import (
	"fmt"
	"time"

	fanoutpattern "github.com/mukezhz/learn/channel/fan_out_pattern"
)

func main() {
	pubSub := fanoutpattern.NewDispatcher()

	for i := 0; i < 3; i++ {
		subscriber := fanoutpattern.NewSubscriber(fmt.Sprintf("Subscriber %d", i))
		pubSub.Subscribe(subscriber)
	}

	producer := fanoutpattern.NewProducer(pubSub)
	go producer.Produce()

	time.Sleep(10 * time.Second)
	pubSub.Close()
}
