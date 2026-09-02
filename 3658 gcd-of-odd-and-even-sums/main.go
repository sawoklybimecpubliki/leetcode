package main

import "fmt"

func main() {
	n := 5
	sumOdd, sumEven, start := 0, 0, 1
	for i := 0; i < n; i++ {
		sumOdd += start
		sumEven += start + 1
		start += 2
	}
	fmt.Println(GCD(sumOdd, sumEven))
}

func GCD(a, b int) int {
	var div int
	if a > b {
		div = b
	} else {
		div = a
	}
	for a%div != 0 || b%div != 0 {
		div--
	}
	return div
}
