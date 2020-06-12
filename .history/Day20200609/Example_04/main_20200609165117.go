package main

import (
	"fmt"
	"time"
)

func main() {
	chananel := make(chan bool)

	go func() {
		time.Sleep(2 * time.Second)
		fmt.Println("2222")
		chananel <- true
	}()

	<-chananel

	fmt.Println("1111")

	
}
