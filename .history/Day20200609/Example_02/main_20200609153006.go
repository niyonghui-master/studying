package main

import "sync"

func main() {

}

func getNumber() int {
	var i int
	var wg sync.WaitGroup
	wg.Add(1)
	go func() {
		i= 5
		wg.Done()
	}()
	wg.Wait()
}