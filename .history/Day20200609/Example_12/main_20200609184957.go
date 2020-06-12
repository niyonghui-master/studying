package main

import "fmt"

type People interface {
	Show()
}

type Student struct{}

func (stu *Student) Show() {

}

func live() People {
	var stu *Student
	return stu
}

func main() {
	if live() == nil {
		fmt.Println("A")
	} else {
		fmt.Println("B", live())
	}

	var m map[string]int

	if m == nil {
		fmt.Println("C")
	}

	var p People
	if p == nil {
		fmt.Println("D")
	}

	

	fmt.Println(deferAndReturn())
}

func deferAndReturn() (a int{
	var a int

	defer func ()  {
		a = 11
	}()

	return
}
