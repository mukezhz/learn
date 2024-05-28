package main

import (
	"log"
	"strconv"
	"time"

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

	streamConfig := &nats.StreamConfig{
		Name:     "ORDERS",
		Subjects: []string{"order.*"},
	}
	_, err = js.AddStream(streamConfig)
	if err != nil {
		log.Fatalf("Stream creation failed: %v", err)
	}

	subject := "order.created"
	for i := 0; ; i++ {
		_, err = js.Publish(subject, []byte("Order message #"+strconv.Itoa(i)))
		if err != nil {
			log.Printf("Error publishing message: %v\n", err)
		} else {
			log.Printf("Published message %d", i)
		}
		time.Sleep(2 * time.Second)
	}
}
