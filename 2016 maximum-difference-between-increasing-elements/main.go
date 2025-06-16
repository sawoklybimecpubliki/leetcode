package main

import "fmt"

func maximumDifference(nums []int) int {
	max_dif := -1
	for i := 0; i < len(nums)-1; i++ {
		for j := i + 1; j < len(nums); j++ {
			if nums[j]-nums[i] > max_dif {
				max_dif = nums[j] - nums[i]
			}
		}
	}
	if max_dif <= 0 {
		return -1
	}
	return max_dif
}

func main() {
	fmt.Println(maximumDifference([]int{1, 5, 2, 10}))
}
