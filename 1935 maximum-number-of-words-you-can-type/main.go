package main

import (
	"fmt"
	"strings"
)

func canBeTypedWords(text string, brokenLetters string) int {

	t := strings.Split(text, " ")
	result := len(t)
	for _, word := range t {
		for _, letter := range brokenLetters {
			if strings.Contains(word, string(letter)) {
				result--
				break
			}
		}
	}
	return result
}

func main() {
	fmt.Println(canBeTypedWords("leet code", "a"))
}
