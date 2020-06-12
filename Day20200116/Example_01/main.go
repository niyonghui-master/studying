package main

import "fmt"

func main() {
	a := make([]int, 5)

	fmt.Println(a)

	var b []int
	b = *new([]int)
	if b == nil {
		fmt.Printf("TTTT, %v\n", b)
	}
	b = append(b, 1)
	fmt.Println(b)

}
