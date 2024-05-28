package main

import (
	"log"

	"github.com/nats-io/nats.go"
)

func main() {
	nc, err := nats.Connect(nats.DefaultURL)
	if err != nil {
		log.Fatal(err)
	}
	defer nc.Close()

	js, err := nc.JetStream()
	if err != nil {
		log.Fatal(err)
	}

	sub, err := js.Subscribe("order.*", func(msg *nats.Msg) {
		log.Printf("Received message: %s", string(msg.Data))
		msg.Ack()
	}, nats.Durable("order-consumer"), nats.ManualAck())
	if err != nil {
		log.Fatalf("Subscribe failed: %v", err)
	}

	log.Println("Listening for messages...", sub.Subject)
	// Keep the subscription running
	select {}
}
