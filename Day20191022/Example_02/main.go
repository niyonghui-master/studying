package main

import "fmt"

func change(s ...string) {
	s[0] = "seekload.net"
	fmt.Printf("%T\n",s)
	fmt.Println(s)
}

func main() {
	arr := [3]string{"Hello","World","Go"}
	change(arr[:]...)
	fmt.Printf("%T\n",arr)
	fmt.Println(arr)
}
