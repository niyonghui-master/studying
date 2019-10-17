package main

import "fmt"

const (
	name = "name"
	c = iota
	d = iota
)

const (
	a = iota
	b = iota
)

func main()  {
	fmt.Println(a)
	fmt.Println(b)
	fmt.Println(c)
	fmt.Println(d)
}
