package main

import "fmt"

func maxFrequencyElements(nums []int) int {
	m := make(map[int]int)
	maxF := 0
	for _, n := range nums {
		m[n]++
		if maxF < m[n] {
			maxF = m[n]
		}
	}
	c := 0
	for _, n := range m {
		if n == maxF {
			c += n
		}
	}
	return c
}

func main() {
	fmt.Println(maxFrequencyElements([]int{1, 2, 5, 3, 4}))
}
