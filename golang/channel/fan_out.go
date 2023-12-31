package main

import (
	"fmt"
	"sync"
)

type customer struct {
	id   int
	name string
}

type data struct {
	msg           string
	consumers     []customer
	consumerCount int
}

// https://stackoverflow.com/questions/16930251/go-one-producer-many-consumers
func main() {
	ch := make(chan *data)
	var wg sync.WaitGroup
	customers := []customer{
		{1, "consumer1"},
		{2, "consumer2"},
		// {3, "consumer3"},
	}
	customerCount := len(customers)

	producer := func() {
		obj := &data{
			msg:           "hello everyone!",
			consumers:     customers,
			consumerCount: customerCount,
		}
		ch <- obj
	}

	go producer()
	for i := 1; i <= customerCount; i++ {
		wg.Add(1)
		go consumer(i, ch, &wg)
	}

	wg.Wait()
}

func consumer(idx int, ch chan *data, wg *sync.WaitGroup) {
	defer wg.Done()
	obj := <-ch
	fmt.Printf("consumer %d received data %v\n", idx, obj.consumers)
	obj.consumerCount--
	if obj.consumerCount > 0 {
		ch <- obj
	} else {
		fmt.Printf("last receiver: %d\n", idx)
	}
}
