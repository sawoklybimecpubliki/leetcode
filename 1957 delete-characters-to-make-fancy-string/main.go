package main

import (
	"fmt"
)

func makeFancyString(s string) string {
	var out []byte
	c := 0
	prev := s[0]
	for i := 0; i < len(s); i++ {
		if prev == s[i] {
			c++
		} else {
			c = 1
			prev = s[i]
		}
		if c <= 2 {
			out = append(out, s[i])
		}
	}
	return string(out)
}

func main() {
	fmt.Println(makeFancyString("leeetcode"))
}
