package main

import "fmt"

func main() {
	var a = []int{1, 2, 3, 4, 5}

	var b = a

	a[0]++

	fmt.Println(a, b)
	fmt.Printf("%p %p\n", a, b)

	var m = map[int]string{1: "a", 2: "b", 3: "c"}
	var m1 = m

	m1[1] = "d"

	fmt.Println(m, m1)
	fmt.Printf("%p %p", m, m1)
}
