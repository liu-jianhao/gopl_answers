package main

import "fmt"

func main() {
	fmt.Println(max(1, 2, 3, 4, 5))
	fmt.Println(min(1, 2, 3, 4, 5))
}

func max(nums ...int) int {
	if len(nums) == 0 {
		panic("at least one element")
	}

	n0 := nums[0]
	for i, n := range nums {
		if i != 0 && n > n0 {
			n0 = n
		}
	}
	return n0
}

func min(nums ...int) int {
	if len(nums) == 0 {
		panic("at least one element")
	}

	n0 := nums[0]
	for i, n := range nums {
		if i != 0 && n < n0 {
			n0 = n
		}
	}
	return n0
}
