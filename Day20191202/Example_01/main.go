package main

import "fmt"

func main() {
	var a = []int{1, 2, 3, 4, 5, 6}

	for _i, _v := range a {
		fmt.Println("key:", _i, "value:", _v, "len:", len(a))

		if _i == 2 {
			a = append(a[0:_i], a[_i+1:]...)
		}

		fmt.Println("a:", a)
	}

	fmt.Println("A:", a)
}
