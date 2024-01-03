package main

import (
	"fmt"
	"time"

	pubsub "github.com/mukezhz/learn/channel/pub_sub"
)

func main() {
	pubSub := pubsub.NewDispatcher()

	for i := 0; i < 3; i++ {
		subscriber := pubsub.NewSubscriber(fmt.Sprintf("Subscriber %d", i))
		pubSub.Subscribe(subscriber)
	}

	producer := pubsub.NewProducer(pubSub)
	go producer.Produce()

	time.Sleep(10 * time.Second)
	pubSub.Close()
}
