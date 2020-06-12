package main

import (
	"fmt"
	"strings"
)

func main() {
	var slice01 = []string{"hello", "liqiang.io"}
	var slice02 = slice01
	fmt.Println("slice02 = ", strings.Join(slice02, " "))
	slice01[0] = "hi"
	fmt.Println("slice02 = ", strings.Join(slice02, " "))


}