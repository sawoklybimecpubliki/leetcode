package main

import "fmt"

func countSubarrays(nums []int) int {
	var result int
	for i := 0; i < len(nums)-2; i++ {
		if float32(nums[i]+nums[i+2]) == float32(nums[i+1])/2 {
			result++
		}
	}
	return result
}

func main() {
	fmt.Println(countSubarrays([]int{1, 5, 1}))
}
