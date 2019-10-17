package main

import (
	"fmt"
	"strings"
)

func main() {
	str := "a.b.c.png"

	strs := strings.Split(str, ".")

	fmt.Println("strs:", strs, strs[len(strs) - 1])
}
