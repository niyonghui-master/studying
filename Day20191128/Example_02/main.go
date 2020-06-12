package main

import (
	"fmt"
	"time"
)

func main() {
	ch := make(chan int, 3)
	v := 0

	// Consumer
	go func() {
		for {
			fmt.Printf("Consumer: %d\n", <-ch)
		}
	}()

	// Producer
	for {
		v++
		fmt.Printf("Producer: %d\n", v)
		ch <- v
		time.Sleep(time.Second)
	}
}
