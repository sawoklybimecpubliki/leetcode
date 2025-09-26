package main

import (
	"fmt"
	"sort"
)

func triangleNumber(nums []int) int {
	sort.Ints(nums)
	totalTriangles := 0

	for longest := len(nums) - 1; longest >= 2; longest-- {
		left, right := 0, longest-1
		for left < right {
			if nums[left]+nums[right] > nums[longest] {
				totalTriangles += (right - left)
				right--
			} else {
				left++
			}
		}
	}
	return totalTriangles
}

func main() {
	fmt.Println(triangleNumber([]int{2, 2, 3, 4}))
}
