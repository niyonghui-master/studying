package main

import (
	"fmt"
	"os"

	"golang.org/x/net/context"
)

func main() {

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
	}
}
