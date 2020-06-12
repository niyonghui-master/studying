package main

import (
	"fmt"
	"strconv"
)

// atoi.go
// ErrRange 表示值超出范围 var ErrRange = errors.New("value out of range")
// ErrSyntax 表示语法不正确 var ErrSyntax = errors.New("invalid syntax")
// NumError 记录转换失败
//type NumError struct {
//	Func string // 失败的函数名(ParseBool, ParseInt, ParseUint, ParseFloat)
//	Num string // 输入的值
//	Err error // 失败的原因（ErrRange，ErrSyntax）
//}
//// int 或 uint 类型的长度（32 或 64）
//const IntSize  = initSize
//const intSize  = 32 << uint(^uint(0)>>63)
//// 实现 Error 接口，输出错误信息
//func (e *NumError) Error() string

func main() {
	// ParseInt 将字符串转换为 int 类型
	// s: 要转换的字符串
	// base: 进位制（2 进制到 36 进制）
	// bitSize: 指定整数类型（0：int, 8 int8, 16: int16, 32 int32 64: int64）
	// 返回转换后的结果和转换时遇到的错误
	// 如果 base 为 0， 则根据字符串的前缀判断进位制（0x: 16, 0:8, 其他：10）
	// func ParseInt(s string, base int, bitSize int)(i int64, err error)
	fmt.Println(strconv.ParseInt("123", 10, 8))
	fmt.Println(strconv.ParseInt("12345", 10, 8))
	fmt.Println(strconv.ParseInt("2147483647", 10, 0))
	fmt.Println(strconv.ParseInt("0xFF", 16, 0))
	fmt.Println(strconv.ParseInt("FF", 16, 0))
	fmt.Println(strconv.ParseInt("0xFF", 0, 0))

	// ParseUint 功能通 ParseInt 一样，只不过返回 uint 类型函数
	// ParseUint(s string, base int, bitSize int)(n uint64, err error)
	fmt.Println(strconv.ParseUint("FF", 16, 8))

	// Atoi 相当于 ParseInt(s, 10, 0)
	// 通常使用这个函数，而不使用 ParseInt
	// func Atoi(s string) (i int, err error)
	fmt.Println(strconv.Atoi("9223372036854775807"))
	fmt.Println(strconv.Atoi("9223372036854775808"))

}
