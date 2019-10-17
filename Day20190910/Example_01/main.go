package main

import "fmt"

func increaseA() int {
	var i int
	defer func() {
		i++
	}()
	return i
}

func increaseB() (r int) {
	defer func() {
		r++
	}()
	return r
}

func main() {
	a := increaseA()
	fmt.Println("A:", a)

	b := increaseB()
	fmt.Println("B:", b)
}
