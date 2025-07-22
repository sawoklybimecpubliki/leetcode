package main

import "fmt"

func maximumUniqueSubarray(nums []int) int {
	result := nums[0]
	start := 0
	windowSum := 0
	inWindow := make(map[int]bool)
	for end := range nums {
		for inWindow[nums[end]] {
			inWindow[nums[start]] = false
			windowSum -= nums[start]
			start++
		}
		windowSum += nums[end]
		inWindow[nums[end]] = true
		result = max(result, windowSum)
	}
	return result
}

func main() {
	fmt.Println(maximumUniqueSubarray([]int{5, 2, 1, 2, 5, 2, 1, 2, 5}))
}
