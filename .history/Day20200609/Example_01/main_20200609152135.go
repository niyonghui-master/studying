package main

import "fmt"

func main() {
	fmt.Println(get)
}

func getNumber() int {
	var i int
	go func() {
		i = 5
	}()
	return i
}
