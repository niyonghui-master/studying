package main

import (
	"runtime"
	"sync"
)

func main() {
	runtime.GOMAXPROCS(1)
	wg := sync.WaitGroup
	
}