package main

import "fmt"

func maxAdjacentDistance(nums []int) int {
	max_diff := ans(nums[0], nums[len(nums)-1])
	for i := 0; i < len(nums)-1; i++ {
		if res := ans(nums[i], nums[i+1]); max_diff < res {
			max_diff = res
		}
	}
	return max_diff
}

func ans(a, b int) int {
	if a > b {
		return a - b
	}
	return b - a
}

func main() {
	fmt.Println(maxAdjacentDistance([]int{1, 2, 4}))
}
