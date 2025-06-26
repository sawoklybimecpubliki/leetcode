package main

import "fmt"

func longestSubsequence(s string, k int) int {
	n := len(s)
	count := 0
	power := 0
	val := int64(0)

	for i := n - 1; i >= 0; i-- {
		if s[i] == '0' {
			count++
		} else {
			if power < 32 {
				val += 1 << power
				if val <= int64(k) {
					count++
				}
			}
		}
		power++
	}
	return count
}

func main() {
	fmt.Println(longestSubsequence("1001010", 5))
}
