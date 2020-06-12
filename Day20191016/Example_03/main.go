package main

import "fmt"

func main() {
	t := 109766 // 毫秒

	s := (t % (60 * 1000)) / 1000
	fmt.Println(s)

	m := (t % (60 * 60 * 1000)) / (60 * 1000)
	fmt.Println(m)

	ms := (t - s*1000 - m*60*1000)

	fmt.Println(ms)
}
