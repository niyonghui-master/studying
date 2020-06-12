package main

type channel chan int

func NoDefer() {
	ch1 := make(channel, 10)
	close(ch1)
}

func Defer() {
	
}