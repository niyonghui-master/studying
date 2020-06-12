package main

import "fmt"

func twoSum(nums []int, target int) []int {
	for k, v := range nums {
		nextV := target - v
		for _k, _v := range nums[k+1:] {
			if _v == nextV {
				return []int{k, _k+k+1}
			}
		}
	}

	return []int{}
}

func main()  {
	nums := []int{4,2,4}
	target := 6
	fmt.Println(twoSum(nums, target))
}
