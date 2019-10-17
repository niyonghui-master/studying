package main

import (
	"fmt"
	"os"
)

func main() {
	file, err := os.Create("./Tmp/test.txt")
	if err != nil {
		fmt.Println("Err:", err)
		return
	}
	defer file.Close()

	i, err := file.WriteString("aaaaa")
	if err != nil {
		return
	}

	fmt.Println("i:", i)
}
