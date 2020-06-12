package main

type People interface {
	Show()
}

type Student struct {}

func (stu *Student) Show() {

}

func 