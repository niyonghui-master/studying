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

	sm1 := struct {
		age int
		m   map[string]string
	}{age: 11, m: map[string]string{"a": "1"}}

	sm2 := struct {
		age int
		m   map[string]string
	}{age: 11, m: map[string]string{"a": "1"}}

	if sm1 == sm2 {
		fmt.Println("sm1 == sm2")
	}
}

/*
	编译不能通过 

	结构体比较：
		1. 结构体只能比较是否相等，不能比较大小。
		2. 
*/ 
