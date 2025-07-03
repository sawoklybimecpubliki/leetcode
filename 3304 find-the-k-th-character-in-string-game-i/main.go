package main

import "fmt"

func kthCharacter(k int) byte {
	word := []byte{'a'}
	for len(word) < k {
		for _, letter := range word {
			word = append(word, letter+1)
		}
	}
	return word[k-1]
}

func main() {
	fmt.Println(string(kthCharacter(10)))
}
