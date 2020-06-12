package main

import "fmt"

func a(a []int) {
	a[0] = 10
	a = append(a, 1)
}

func main() {
	A := []int{1}
	a(A)

	fmt.Println(A)
}
