package main

import "fmt"

func maxOperations(s string) int {
	ops := 0
	zeroRuns := 0
	prev := '1'

	for i := len(s) - 1; i >= 0; i-- {
		c := s[i]
		if c == '0' {
			if prev == '1' {
				zeroRuns++
			}
			prev = '0'
		} else { // '1'
			ops += zeroRuns
			prev = '1'
		}
	}
	return ops
}

func main() {
	fmt.Println(maxOperations("110"))
}
