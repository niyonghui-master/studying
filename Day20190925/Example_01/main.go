package main

import (
	"log"
	"time"
)

func main() {
	s := make(chan int, 5)

	go fbqn(s)

	for {
		log.Println(<-s)
		time.Sleep(5 * time.Second)
	}
}

func fbqn(s chan int) {
	x, y := 0, 1
	for {
		if 1 > len(s) {
			s <- x
			x, y = y, x+y
		}
	}
}
