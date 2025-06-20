package main

import "fmt"

func getDigitSum(x int) int {
	sum := 0
	for x > 0 {
		sum += x % 10
		x /= 10
	}
	return sum
}

func countLargestGroup(n int) int {
	cnt := make([]int, 37)

	for i := 1; i <= n; i++ {
		cnt[getDigitSum(i)]++
	}

	maxFreq, res := 0, 0
	for _, freq := range cnt {
		if freq > maxFreq {
			maxFreq, res = freq, 1
		} else if freq == maxFreq {
			res++
		}
	}

	return res
}

func main() {
	fmt.Println(countLargestGroup(13))
}
