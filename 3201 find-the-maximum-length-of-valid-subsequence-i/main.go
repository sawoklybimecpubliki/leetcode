package main

import "fmt"
import "math"

func maximumLength(nums []int) int {
	if len(nums) == 0 {
		return 0
	}

	oddCount, evenCount, alternating := 0, 0, 0
	expectOdd := nums[0]%2 == 1

	for _, n := range nums {
		isOdd := n%2 == 1

		if isOdd == expectOdd {
			alternating++
			expectOdd = !expectOdd
		}

		if isOdd {
			oddCount++
		} else {
			evenCount++
		}
	}

	return int(math.Max(float64(alternating), math.Max(float64(oddCount), float64(evenCount))))
}

func main() {
	fmt.Println(maximumLength([]int{1, 2, 1, 1, 2, 1, 2}))
}
