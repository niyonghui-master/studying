package main

import (
	"fmt"
	"os"
	"time"

	"golang.org/x/net/context"
)

func main() {
	timer := time.NewTimer(time.Second * 5)
	data := []int{1, 2, 3, 10, 999, 8, 345, 7, 98, 33, 66, 77, 88, 68, 96}
	dataLen := len()
}

func SearchTarget(ctx context.Context, data []int, target int, resultChan chan bool) {
	for _, v := range data {
		select {
		case <-ctx.Done():
			fmt.Fprintf(os.Stdout, "Task cancelded! \n")
			return
		default:
		}
		// 模拟一个耗时查找，这里只是比对值，真实开发中可以是其他操作
		fmt.Fprintf(os.Stdout, "v: %d \n", v)
		time.Sleep(time.Millisecond * 1500)
		if target == v {
			resultChan <- true
			return
		}
	}
}
