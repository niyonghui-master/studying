package main

import "fmt"

func main() {
	defer func() {
		fmt.Println("First")
		fmt.Println(recover())
	}()

	defer func() {
		fmt.Println("Second -")
		defer fmt.Println(recover())
		fmt.Println("Second --")
		panic(1)
	}()
	defer recover()

	fmt.Println("Main")

	panic(2)
}
