package main

import (
	"fmt"
	"strconv"
)

func main() {
	// FormatInt 将int 型整数 i 转换为字符串形式
	// base: 进位制（2 进制到 36 进制）
	// 大于 10 进制的数，返回值使用小写字母 'a' 到 'z'
	// func FormatInt(i int64, base int) string
	i := int64(-2048)
	fmt.Println(strconv.FormatInt(i, 2))  // -100000000000
	fmt.Println(strconv.FormatInt(i, 8))  // -4000
	fmt.Println(strconv.FormatInt(i, 10)) // -2048
	fmt.Println(strconv.FormatInt(i, 16)) // -2048

	// Itoa 相当于 FormatInt(i, 10)
	// func Itoa(i int) string

	fmt.Println(strconv.Quote(`C:\windons`))
	fmt.Println(strconv.QuoteToASCII("Hello 世界！"))
	fmt.Println(strconv.QuoteRune('界'))
}
