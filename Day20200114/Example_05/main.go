package main

import "fmt"

func main() {
	numbers := [9]int{0,1,2,3,4,5,6,7,8}

	numbers1 := numbers[0:5]

	numbers2 := numbers1[2:4]
	numbers2[0] = 10

	fmt.Println(numbers)
	fmt.Println(numbers1)
	fmt.Println(numbers2)


}