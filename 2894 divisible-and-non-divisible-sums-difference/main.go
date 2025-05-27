package main

import "fmt"

func differenceOfSums(n int, m int) int {
	sum1, sum2 := 0, 0
	for i := range n + 1 {
		if i%m == 0 {
			sum2 += i
		} else {
			sum1 += i
		}
	}
	return sum1 - sum2
}

func main() {
	fmt.Println(differenceOfSums(5, 1))
}
