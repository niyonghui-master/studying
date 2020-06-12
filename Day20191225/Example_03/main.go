package main

import (
	"log"
	"time"
)

func main() {
	var a map[string]int
	a = make(map[string]int)

	a["aaa"] = 1

	go func() {
		time.Sleep(2)
		log.Println(a)
	}()

	a = make(map[string]int)

	log.Println("1111", a)

	time.Sleep(5)
}
