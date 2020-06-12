package main

import "fmt"

type People struct {}

func (p *People) ShowA() {
	fmt.Println("show A")
	p.ShowB()
}

func (p *People) ShowB() {
	fmt.Println("show B")
}