package main

import "fmt"

func main() {
	chananel := make(chan bool, 2)

	<- chananel 

	fmt.Println()
}