package main

import (
	"fmt"
	"sort"
)

func main() {
	count := make(map[int]int)
	nums := []int{3, 2, 3, 4, 2, 3, 1, 5, 5}
	for _, num := range nums {
		count[num]++
		if count[num] > len(nums)/2 {
			fmt.Println(num)
			return
		}
	}
	sort.Ints(nums)
	fmt.Println(nums)
}
