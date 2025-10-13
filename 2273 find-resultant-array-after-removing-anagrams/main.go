package main

import "fmt"

func removeAnagrams(words []string) []string {
	for i := 1; i < len(words); {
		if len(words[i]) == len(words[i-1]) && isAnagram(words[i], words[i-1]) {
			words = append(words[:i], words[i+1:]...)
		} else {
			i++
		}
	}
	return words
}

func isAnagram(word1, word2 string) bool {
	m := make(map[int32]int)
	for _, l := range word1 {
		m[l]++
	}
	for _, l := range word2 {
		if v, ok := m[l]; ok && v > 0 {
			m[l]--
		} else {
			return false
		}
	}
	return true
}

func main() {
	fmt.Println(removeAnagrams([]string{"a", "az"}))
}
