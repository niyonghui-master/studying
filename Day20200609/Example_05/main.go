package main

import "fmt"

type student struct {
	Name string
	Age  int
}

func main() {
	m := make(map[string]*student)
	stus := []student{
		{Name: "A", Age: 24},
		{Name: "B", Age: 23},
		{Name: "C", Age: 22},
	}

	for _, stu := range stus {
		m[stu.Name] = &stu
	}

	fmt.Println(m)
}
