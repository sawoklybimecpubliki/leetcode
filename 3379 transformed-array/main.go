package main

import "fmt"

func constructTransformedArray(nums []int) []int {
	ln := len(nums)
	res := make([]int, ln)
	for i, n := range nums {
		idx := (i + n%ln + ln) % ln
		res[i] = nums[idx]
	}
	return res
}

func main() {
	fmt.Println(constructTransformedArray([]int{3, -2, 1, 1}))
}
