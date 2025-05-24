package main

import (
	"fmt"
)

func findWordsContaining(words []string, x byte) []int {
	var out []int
	for i, word := range words {
		for _, letter := range word {
			if x == byte(letter) {
				out = append(out, i)
				break
			}
		}
		/*if strings.Contains(word, string(x)) {
			out = append(out, i)
		}*/
	}
	return out
}

func main() {
	fmt.Println(findWordsContaining([]string{"abc", "bcd", "aaaa", "cbc"}, 'a'))
}
