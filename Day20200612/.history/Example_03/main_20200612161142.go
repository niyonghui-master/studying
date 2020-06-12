package main

import "fmt"

func main() {
	list := new([]int)
	list = append(list, 1)
	fmt.Println(list)
}

/*
	不能编译通过，new([]int) 之后的list是一个 *[]int 类型的指针，不能对指针执行 append 操作。可以使用 make
*/ 