package main

import "fmt"

type Test interface {}

type Test1 interface {
	TestFunc()
}

type Structure struct {
	a int
}

func (s *Structure) TestFunc() {
	fmt.Println("OK, Let's rock and roll!")
}

func fTest(t test)