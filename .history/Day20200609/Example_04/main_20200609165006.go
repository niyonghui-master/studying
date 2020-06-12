package main

import (
	"fmt"
	"time"
)

func main() {
	chananel := make(chan bool)

	go func() {
		chananel 
	}()

	<- chananel 

	fmt.Println("1111")

	time.Sleep(2 * time.Second)
}