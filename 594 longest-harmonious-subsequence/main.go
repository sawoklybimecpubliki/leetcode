package main

import (
	"fmt"
	"sort"
)

func findLHS(nums []int) int {
	sort.Ints(nums)
	l, r, ans := 0, 0, 0
	n := len(nums)

	for r < n {
		diff := nums[r] - nums[l]
		if diff == 1 {
			if r-l+1 > ans {
				ans = r - l + 1
			}
		}
		if diff <= 1 {
			r++
		} else {
			l++
		}
	}

	return ans
}

func main() {
	fmt.Println(findLHS([]int{3, 2, 1}))
}
