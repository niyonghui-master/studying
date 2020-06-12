package main

import "fmt"

type People struct {}

func (p *People) ShowA() {
	fmt.Println("show ")
}