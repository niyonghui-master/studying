package main

func main() {
	c := make(chan int, 10)
	// 从一个永远都不可能有值的通道中读取数据，会发生死锁，因为会阻塞主程序的执行， 如果是在新开的协程中是没有问题的
	//<-c

	go func() {
		<-c
	}()
}
