package main

import (
	"fmt"
	"strconv"
)

func sequentialDigits(low int, high int) []int {
	digits := "123456789"

	ans := []int{}

	minLen := len([]byte(fmt.Sprintf("%d", low)))
	maxLen := len([]byte(fmt.Sprintf("%d", high)))

	for length := minLen; length <= maxLen; length++ {

		for start := 0; start+length <= 9; start++ {

			num, _ := strconv.Atoi(digits[start : start+length])

			if num >= low && num <= high {
				ans = append(ans, num)
			}
		}
	}

	return ans
}

func main() {
	fmt.Println(sequentialDigits(100, 300))
}
