package main

import "fmt"

type Tester interface {
	getName() string
}

type Tester2 interface {
	printName() 
}

type Person struct {
	name string
}

func (p Person) getName() string {
	return p.name
}

func (p Person) printName() {
	fmt.Println(p.name)
}

func m