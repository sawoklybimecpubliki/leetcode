package main

import (
	"fmt"
	"sort"
)

func minimizeMax(nums []int, p int) int {
	sort.Ints(nums)
	left, right := 0, nums[len(nums)-1]-nums[0]
	for left < right {
		mid := (left + right) / 2
		if countPairs(nums, mid, p) {
			right = mid
		} else {
			left = mid + 1
		}
	}
	return right
}

func countPairs(nums []int, limit, p int) bool {
	for i := 1; i < len(nums) && p > 0; i++ {
		if nums[i]-nums[i-1] <= limit {
			p--
			i++
		}
	}
	return p <= 0
}

func main() {
	fmt.Println(minimizeMax([]int{10, 1, 2, 7, 1, 3}, 2))
}
