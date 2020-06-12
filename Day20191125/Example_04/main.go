package main

import "fmt"

func F(n int) func() int {
	return func() int {
		n++
		return n
	}
}

func main() {
	f := F(5)

	fmt.Println(f())
	fmt.Println(f())
	fmt.Println(f())
	fmt.Println(f())
}
