package main

import (
	"fmt"
	"time"
)

func main() {
	chananel := make(chan bool, 2)

	<- chananel 

	fmt.Println("1111")

	time.Sleep(2 * time.Second)
}