package main

import "fmt"

func threeConsecutiveOdds(arr []int) bool {
	c := 0
	for _, i := range arr {
		if i%2 != 0 {
			c++
		} else {
			c = 0
		}
		if c == 3 {
			return true
		}
	}
	return false
}

func main() {
	fmt.Println(threeConsecutiveOdds([]int{1, 2, 34, 3, 4, 5, 7, 23, 12}))
}
