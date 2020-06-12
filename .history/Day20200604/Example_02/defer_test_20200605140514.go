package main

import "testing"

type channel chan int

func NoDefer() {
	ch1 := make(channel, 10)
	close(ch1)
}

func Defer() {
	ch2 := make(channel, 10)
	defer close(ch2)
}

func BenchmarkNoDefer(b *testing.B) {
	for 
}