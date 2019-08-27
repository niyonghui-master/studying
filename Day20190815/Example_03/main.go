package main

import "fmt"

func main() {
	c := make(chan int)
	// 主程序往一个没有消费者的通道中写入数据时会发生死锁，因为会阻塞主程序的执行
	c <- 1

	fmt.Println("Test...")
}
