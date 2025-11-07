package main

import "fmt"

func countValidSelections(nums []int) int {
	suffixSum := 0
	for _, num := range nums {
		suffixSum += num
	}

	prefixSum := 0
	result := 0
	for _, num := range nums {
		suffixSum -= num
		prefixSum += num
		diff := abs(prefixSum - suffixSum)
		if num == 0 && diff < 2 {
			result += 2 - diff
		}
	}
	return result
}

func abs(x int) int {
	if x < 0 {
		return -x
	}
	return x
}

func main() {
	fmt.Println(countValidSelections([]int{1, 0, 2, 0, 3}))
}
