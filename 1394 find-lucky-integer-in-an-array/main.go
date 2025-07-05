package main

import "fmt"

func findLucky(arr []int) int {
	m := make(map[int]int)
	maxN := -1
	for _, n := range arr {
		m[n]++
	}
	for key, val := range m {
		if val == key {
			if maxN < val {
				maxN = val
			}
		}
	}
	return maxN
}

func main() {
	fmt.Println(findLucky([]int{1, 2, 2, 3, 3, 3}))
}
