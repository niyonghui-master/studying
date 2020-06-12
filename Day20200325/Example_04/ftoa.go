package main

import (
	"fmt"
	"strconv"
)

func main()  {
	// FormatFloat 将浮点数 f 转换为字符串值
	// f：要转换的浮点数
	// fmt: 格式标记 （b、e、E、f、g、G）
	// prec：精度（数字部分的长度，不包括指数部分）
	// bitSize：指定浮点数类型（32：float32 64: float64）
	// 格式标记
	// 'b'(-ddddp±ddd，二进制指数)
	// 'e'(-d.dddde±dd，十进制指数)
	// 'E'(-d.ddddE±dd，十进制指数)
	// 'f'(-ddd.dddd，没有指数)
	// 'g'('e':大指数，'f':其它情况)
	// 'G'('E':大指数，'f':其它情况)
	// func FormatFloat(f float64, fmt byte, prec, bitSize int) string
	f := 100.12345678901234567890123456789
	fmt.Println(strconv.FormatFloat(f, 'b', 5, 32))	// 13123382p-17
	fmt.Println(strconv.FormatFloat(f, 'e', 5, 32))	// 1.00123e+02
}
