package main

import (
	"fmt"
	"time"
)

func producer(ch chan<- string) {
	for i := 0; i < 10; i++ {
		ch <- fmt.Sprintf("Message %d", i)
		time.Sleep(time.Second)
	}
}

func consumer(ch <-chan string) {
	for msg := range ch {
		fmt.Println("Received:", msg)
	}
}

func main() {
	ch := make(chan string)

	go producer(ch)
	go consumer(ch)

	time.Sleep(10 * time.Second)
}
