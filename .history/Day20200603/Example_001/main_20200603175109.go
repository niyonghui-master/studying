package main

import "fmt"

type S struct{}

func main() {
    var x S
    y := &x
	w := identity(y)
	
	fmt.Println(w)
}

func identity(z *S) *S {
    return z
}