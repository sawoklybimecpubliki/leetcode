package main

import "fmt"

func maxDifference(s string) int {
	count := make(map[int32]int)
	for _, letter := range s {
		count[letter]++
	}
	max_odd, min_even := 0, 100
	for _, val := range count {
		if val%2 == 0 {
			if min_even > val {
				min_even = val
			}
		} else {
			if max_odd < val {
				max_odd = val
			}
		}
	}
	return max_odd - min_even
}

func main() {
	fmt.Println(maxDifference("abcabcab"))
}
