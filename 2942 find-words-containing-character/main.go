package main

import (
	"fmt"
	"strings"
)

func findWordsContaining(words []string, x byte) []int {
	var out []int
	for i, word := range words {
		if strings.Contains(word, string(x)) {
			out = append(out, i)
		}
	}
	return out
}

func main() {
	fmt.Println(findWordsContaining([]string{"abc", "bcd", "aaaa", "cbc"}, 'a'))
}
