package main

import (
	"fmt"
	"time"
)

func fanIn(input1, input2 <-chan string) <-chan string {
	c := make(chan string)
	go func() {
		for {
			select {
			case s := <-input1:
				c <- s
			case s := <-input2:
				c <- s
			}
		}
	}()
	return c
}

func main() {
	input1 := generateString("input1")
	input2 := generateString("input2")
	c := fanIn(input1, input2)
	for i := 0; i < 10; i++ {
		fmt.Println(<-c)
	}
	fmt.Println("Done")
}

func generateString(msg string) <-chan string {
	c := make(chan string)
	go func() {
		for i := 0; ; i++ {
			c <- fmt.Sprintf("Hello %d: %v", i, msg)
			time.Sleep(time.Millisecond * 500)
		}
	}()
	return c
}
