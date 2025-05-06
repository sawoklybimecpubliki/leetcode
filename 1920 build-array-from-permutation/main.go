package main

import "fmt"

func buildArray(nums []int) []int {
	var result []int

	for i := 0; i < len(nums); i++ {
		result = append(result, nums[nums[i]])
	}
	return result
}

func main() {
	fmt.Println(buildArray([]int{5, 0, 1, 2, 3, 4}))
}
