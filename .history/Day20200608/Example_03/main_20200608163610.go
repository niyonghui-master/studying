package main

import (
	"fmt"
	"strings"
)

func main() {
	var slice01 = []string{"hello", "NYH"}
	var slice02 = make([]string, len(slice01))
	copy(slice02, slice01)
	fmt.Println("slice02 = ", strings.Join(slice02, " "))
	slice01[0] = "h1"
	fmt.Println("slice02 = ", strings.Join(slice02, ""))
}