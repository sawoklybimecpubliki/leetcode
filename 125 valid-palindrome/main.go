package main

import (
	"fmt"
	"strings"
)

func isPalindrome(s string) bool {
	if s == " " {
		return true
	}
	s = strings.ToLower(s)
	for i, j := 0, len(s)-1; i <= j; {
		if (s[i] < 'a' || s[i] > 'z') && (s[i] < '0' || s[i] > '9') {
			i++
			continue
		}
		if (s[j] < 'a' || s[j] > 'z') && (s[j] < '0' || s[j] > '9') {
			j--
			continue
		}
		if s[i] != s[j] {
			return false
		}
		i, j = i+1, j-1
	}
	return true
}

func main() {
	fmt.Println(isPalindrome("a."))
}
