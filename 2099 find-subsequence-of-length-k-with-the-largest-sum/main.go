package main

import (
	"fmt"
)

func maxSubsequence(nums []int, k int) []int {
	result := make([]int, k)
	result = nums[:k]
	index := 0
	for i := k; i < len(nums); i++ {
		if index == 0 {
			index = indexMin(result)
		}
		if result[index] < nums[i] {
			result = append(result[:index], result[index+1:]...)
			result = append(result, nums[i])
			index = 0
		}
	}

	return result
}

func indexMin(nums []int) int {
	var index int
	for i := 1; i < len(nums); i++ {
		if nums[index] > nums[i] {
			index = i
		}
	}
	return index
}

func main() {
	fmt.Println(maxSubsequence([]int{3, 4, 3, 3}, 2))
}
