package main

import (
	"fmt"
	"sort"
)

func sortByBits(arr []int) []int {

	lst := make([][]int, 15)

	for _, num := range arr {
		bits := bitCount(num)
		lst[bits] = append(lst[bits], num)
	}

	ans := []int{}

	for _, bucket := range lst {
		sort.Ints(bucket)
		ans = append(ans, bucket...)
	}

	return ans
}

func bitCount(n int) int {
	count := 0
	for n > 0 {
		count += n & 1
		n >>= 1
	}
	return count
}

func main() {
	fmt.Println(sortByBits([]int{1, 2, 3, 4, 5, 6, 7, 8}))
}
