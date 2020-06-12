package main

import (
	"fmt"
	"runtime"
)

func main() {
	runtime.GOMAXPROCS(1)
	int_chan := make(chan int, 1)
	stringChan := make(chan string, 1)
	int_chan <- 1
	stringChan <- "hello"
	select {
	case value := <-int_chan:
		fmt.Println(value)
	case value := <-stringChan:
		panic(value)
	}
}
