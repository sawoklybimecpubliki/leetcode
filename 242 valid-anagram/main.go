package main

import (
	"fmt"
)

func isAnagram(s string, t string) bool {
	letterS := make(map[uint8]int)
	letterT := make(map[uint8]int)
	if len(s) != len(t) {
		return false
	}
	for i := 0; i < len(s); i++ {
		fmt.Println(s[i])
		letterS[s[i]]++
		letterT[t[i]]++
	}
	for i := 0; i < len(s); i++ {
		if letterS[s[i]] != letterT[s[i]] {
			return false
		}
	}
	return true
}

func main() {
	fmt.Println(isAnagram("anagram", "nagaram"))
}
