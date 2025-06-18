package main

import (
	"fmt"
	"sort"
)

func divideArray(nums []int, k int) [][]int {
	var out [][]int
	sort.Ints(nums)
	n := len(nums) / 3
	for i := 0; i < n; i++ {
		for j := i * 3; j < (i+1)*3-2; j++ {
			if nums[j+2]-nums[j] > k {
				return [][]int{}
			}
		}
		out = append(out, nums[i*3:(i+1)*3])
	}
	return out
}

func main() {
	fmt.Println(divideArray([]int{4, 2, 9, 8, 2, 12, 7, 12, 10, 5, 8, 5, 5, 7, 9, 2, 5, 11}, 14))
}
