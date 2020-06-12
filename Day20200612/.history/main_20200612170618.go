package main

import "fmt"

func main() {
	sn1 := struct {
		age  int
		name string
	}{age: 11, name: "NYH"}

	sn2 := struct {
		age  int
		name string
	}{age: 11, name: "NYH"}

	if sn1 == sn2 {
		fmt.Println("sn1 == sn2")
	}

	// sm1 := struct {
	// 	age int
	// 	m   map[string]string
	// }{age: 11, m: map[string]string{"a": "1"}}

	// sm2 := struct {
	// 	age int
	// 	m   map[string]string
	// }{age: 11, m: map[string]string{"a": "1"}}

	// if sm1 == sm2 {
	// 	fmt.Println("sm1 == sm2")
	// }

	newB
}

/*
	编译不能通过

	结构体比较：
		1. 结构体只能比较是否相等，不能比较大小。
		2. 相同类型的结构体才能够进行比较，结构体是否相同不但与属性类型有关，还与属性顺序相关，sn3 与 sn1 就是不同的结构体
		sn3 := struct {
			name string
			age int
		}{age : 11, name : "NYH"}
		3. 如果 struct 的所有成员都可以比较，则该 struct 就可以通过 == 或 != 进行比较是否相等，比较时逐个项进行比较，如果每一项都相等，则两个结构体才相等，否则不想等。

		那什么是可比较的呢，常见的有bool、数值型、字符、指针、数组等，像切片、map、函数等是不能比较的。
*/


/*	
	字符串连接：
		1. str := "abc" + "123"
		2. fmt.Sprintf("abc%d", 123)
		3. strings.Join()  -> func Join(elems []string, sep string) string
		4. buffer.WriterString() -> func (*Buffer) WriteString
*/ 