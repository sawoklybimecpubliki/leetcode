package main

import (
	"fmt"
)

func main() {
	s := "anagram"
	t := "nagaram"
	letterS := make(map[uint8]int)
	letterT := make(map[uint8]int)
	if len(s) != len(t) {
		fmt.Printf("len stirngs does not match")
		return
	}
	for i := 0; i < len(s); i++ {
		fmt.Println(s[i])
		letterS[s[i]]++
		letterT[t[i]]++
	}
	for i := 0; i < len(s); i++ {
		if letterS[s[i]] != letterT[s[i]] {
			fmt.Printf("letters do not match")
			return
		}
	}
	fmt.Printf("letters match")
	return
}
