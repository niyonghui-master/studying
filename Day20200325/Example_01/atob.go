package main

import (
	"fmt"
	"strconv"
)

// ParseBool 将字符串转换为布尔值
// 它接收真值：1, t, T, TRUE, true, True
// 它接收假值：0, f, F, FALSE, false, False
// 其他任何值都返回一个错误
// func ParseBool(str string) (value bool, err error)

func main() {
	fmt.Println(strconv.ParseBool("1"))
	fmt.Println(strconv.ParseBool("t"))
	fmt.Println(strconv.ParseBool("T"))
	fmt.Println(strconv.ParseBool("true"))
	fmt.Println(strconv.ParseBool("True"))
	fmt.Println(strconv.ParseBool("TRUE"))
	fmt.Println(strconv.ParseBool("TRue"))

	// FormatBool 将布尔值转换为字符串 "true" 或 "false"
	// func FormatBool(b bool) string
	fmt.Println(strconv.FormatBool(0 < 1))

	// AppendBool 将布尔值 b 转换为字符串 "true" 或 "false"
	// 然后将结果追加到 dst 的尾部，返回追加后的 []byte
	// func AppendBool(dst []byte, b bool) []byte

	rst := make([]byte, 0)
	rst = strconv.AppendBool(rst, 0 < 1)
	fmt.Printf("%s\n", rst)
	rst = strconv.AppendBool(rst, 0 > 1)
	fmt.Printf("%s\n", rst)
}
