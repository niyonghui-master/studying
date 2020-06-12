package main

import "fmt"

func main() {
	// 浅拷贝
	a := []int{1, 2, 3, 4, 5, 6}
	fmt.Printf("a地址1: %p\n", a)
	b := a[0:3]
	fmt.Println("a is:", a)
	fmt.Println("b is:", b)

	fmt.Printf("a地址2：%p")
}
