package main

import "fmt"

func triangularSum(nums []int) int {
	n := len(nums)

	for size := n; size > 1; size-- {
		for i := 0; i < size-1; i++ {
			nums[i] = (nums[i] + nums[i+1]) % 10
		}
	}

	return nums[0]
}

func main() {
	fmt.Println(triangularSum([]int{1, 2, 3, 4, 5}))
}
