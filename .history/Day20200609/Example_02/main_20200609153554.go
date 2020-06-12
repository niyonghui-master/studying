package main

import (
	"fmt"
	"sync"
)

func main() {
	fmt.Println(getNumber())
	fmt.Println(getNumberByChan())
}

func getNumber() int {
	var i int
	var wg sync.WaitGroup
	wg.Add(1)
	go func() {
		i = 5
		wg.Done()
	}()
	wg.Wait()
	return i
}

func getNumberByChan() int {
	var i int
	done := make(chan struct{})
	go func() {
		i = 5
		done <- struct{}{}
	}()
	<-done
	return i
}

//  首先，创建一个结构体包含我们想用互斥锁保护的值和一个mutex实例
type SafeNumber struct {
	val int
	m   sync.Mutex
}

func (i *SafeNumber) Get


