package main

import "fmt"

var (
	size     = 1024 // 原文 -> size := 1024 
	max_size = size * 2
)

func main() {
	fmt.Println(size, max_size)
}

/*
	简短声明限制：
		1. 必须使用显示初始化
		2. 不能提供数据类型，编译器会自动推导
		3.只能在函数内部使用简短mo
*/ 
