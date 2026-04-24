package main

import "fmt"

func closestTarget(words []string, target string, startIndex int) int {
	n := len(words)
	k := 1
	for i, j := startIndex, startIndex; k <= (n/2)+1; i, j = i+1, j-1 {
		if target == words[startIndex] {
			return 0
		}
		if target == words[(j-1+n)%n] || target == words[(i+1)%n] {
			return k
		}
		k++
	}
	return -1
}

func main() {
	words := []string{"hello", "i", "am", "leetcode", "hello"}
	target := "leetcode"
	startIndex := 1
	fmt.Println(closestTarget(words, target, startIndex))
}
