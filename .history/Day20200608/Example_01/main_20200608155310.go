package main

import "fmt"

func main() {
	// 浅拷贝
	a := []int{1, 2, 3, 4, 5, 6}
	fmt.Printf("a地址1: %p\n", a)
	b := a
}
