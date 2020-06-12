package main

type People interface {
	Show()
}

type Student struct {}

func (stu *Student) Show() {

}

func live() People {
	var stu *Student
	return stu
}

func main() {
	if live() == nil 
}