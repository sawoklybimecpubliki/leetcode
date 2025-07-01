package main

import "fmt"

func possibleStringCount(word string) int {
	result := 1
	for i := 1; i < len(word); i++ {
		if word[i] == word[i-1] {
			result++
		}
	}
	return result
}

func main() {
	fmt.Println(possibleStringCount("abcd"))
}
