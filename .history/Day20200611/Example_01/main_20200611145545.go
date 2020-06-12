package main

import "fmt"

type student struct {
	Name string
}

func test(v interface{}) {
	switch msg := v.(type) {
	case *student, student:
		fmt.Println(msg.(student).Name)
	}
}
