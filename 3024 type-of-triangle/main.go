package main

import "fmt"

func triangleType(nums []int) string {
	if nums[0] == nums[1] || nums[1] == nums[2] || nums[0] == nums[2] {
		if nums[0] == nums[1] && nums[0] == nums[2] {
			return "equilateral"
		}
		return "isosceles"
	}
	return "scalene"
}

func main() {
	fmt.Println(triangleType([]int{3, 5, 4}))
}
