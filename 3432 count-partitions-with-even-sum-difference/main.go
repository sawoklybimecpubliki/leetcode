package main

import "fmt"

func countPartitions(nums []int) int {
	total := 0
	for _, v := range nums {
		total += v
	}
	left, ans := 0, 0
	for i := 0; i < len(nums)-1; i++ {
		left += nums[i]
		right := total - left
		if ((left - right) & 1) == 0 {
			ans++
		}
	}
	return ans
}

func main() {
	fmt.Println(countPartitions([]int{10, 10, 3, 7, 6}))
}
