package main

import "fmt"

func findNumbers(nums []int) int {
	var result int
	for _, n := range nums {
		switch {
		case n >= 100000:
			result++
		case n >= 1000 && n < 10000:
			result++
		case n >= 10 && n < 100:
			result++
		}
	}
	return result
}

func main() {
	fmt.Println(findNumbers([]int{555, 901, 482, 1771}))
}
