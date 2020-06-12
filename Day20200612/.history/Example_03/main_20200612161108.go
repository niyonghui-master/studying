package main

import "fmt"

func main() {
	list := new([]int)
	list = append(list, 1)
	fmt.Println(list)
}

/*
	不能编译通过，new([]int) 之后的list是一个 *[]int 类型
*/ 