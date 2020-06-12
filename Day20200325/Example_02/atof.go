package main

import (
	"fmt"
	"strconv"
)

func main()  {
	// ParseFloat 将字符串转换为浮点数
	// s: 要转换的字符串
	// bitSize: 指定浮点类型（32：float32、64：float64）
	// 如果 s 是合法的格式，而且接近一个浮点值，
	// 则返回浮点数的四舍五入值（依据 IEEE754 的四舍五入标准）
	// 如果 s 不是合法的格式，则返回 "语法错误"
	// 如果转换结果超出 bitsize 范围，则返回 "超出范围"
	// func ParseFloat(s string, bitSize int) (f float64, err error)
	s := "0.12345678901234567890"
	f, err := strconv.ParseFloat(s, 32)
	fmt.Println(f, err)
	fmt.Println(float32(f), err)
	f, err = strconv.ParseFloat(s, 64)
	fmt.Println(f, err)
}
