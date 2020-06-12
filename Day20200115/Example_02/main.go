package main

import (
	"NYH/Day20200115/Example_02/library"
	"fmt"
)

var lib *library.MusicManager
var id int
var ctrl, signal chan int

func handleLibCommands(tokens []string) {
	switch tokens[1] {
	case "list":
		for i := 0; i < lib.Len(); i++ {
			e, _ := lib.Get(i)
			fmt.Println(i+1, ":", e.Name, e.Artist, e.Source, e.Type)
		}
	case "add":
	}
}
