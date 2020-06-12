package main

import "fmt"

type Summer interface {
	Sum() int
}

type ints []int

func (i *ints) Sum() int {
	s := 0
	for _, v := range i {
		s += v
	}
	return s
}

func main() {
	var i ints
	var s Summer = &i              // nil value can satisfy interface
	fmt.Println(s == nil, s.Sum()) // true, 0
}
