package main

import "fmt"

func merge(out chan<- int, a, b <-chan int) {
	var aClosed, bClosed bool

	for !aClosed || !bClosed {
		select {
		case v, ok := <-a: // 此时通道关闭后，就不会再进行获取了
			if !ok {
				aClosed = true
				fmt.Println("a is closed")
				continue
			}
			out <- v
		case v, ok := <-b:
			if !ok {
				bClosed = true
				fmt.Println("b is closed")
				continue
			}
			out <- v
		}
	}
	close(out) // 需要再不使用后进行close操作
}

func main() {
	var (
		out chan<- int
		a   chan int
		b   chan int
	)

	out = make(chan<- int)
	a = make(<-chan int)
	b = make(<-chan int)

	merge(out, a, b)
}
