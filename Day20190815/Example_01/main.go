package main

import "fmt"

func main() {
	test()
}

func test() {
	defer func() {
		if err := recover(); err != nil {
			fmt.Println("e1:", err)
		}
	}()

	panic("no1")

	// 以下代码不会执行

	fmt.Println("123456")

	defer func() {
		if err := recover(); err != nil {
			fmt.Println("e2:", err)
		}
	}()

	panic("no2")
}
