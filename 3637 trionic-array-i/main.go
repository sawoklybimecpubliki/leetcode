package main

import "fmt"

func isTrionic(nums []int) bool {
	n := len(nums)
	if n < 4 {
		return false
	}

	i := 0

	for i+1 < n && nums[i] < nums[i+1] {
		i++
	}

	if i == 0 || i == n-1 {
		return false
	}

	j := i

	for j+1 < n && nums[j] > nums[j+1] {
		j++
	}

	if j == i || j == n-1 {
		return false
	}

	for j+1 < n && nums[j] < nums[j+1] {
		j++
	}

	return j == n-1
}

func main() {
	fmt.Println(isTrionic([]int{1, 3, 5, 4, 2, 6}))
}
