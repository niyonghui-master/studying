package main

import "fmt"

func main() {
	s := []int{1, 1, 1}
	changeSlice(s)

	fmt.Println(s)
}

func changeSlice(x []int) {
	for i, _ := range x {
		x[i] += 1
	}
}
