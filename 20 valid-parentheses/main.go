package main

import "fmt"

func isValid(s string) bool {
	var expectation []uint8
	for i := 0; i < len(s); i++ {
		if s[i] == '(' || s[i] == '[' || s[i] == '{' {
			expectation = append(expectation, s[i])
		}
		if s[i] == ')' || s[i] == ']' || s[i] == '}' {
			if expectation[len(expectation)-1] == '(' {
				if s[i] != ')' {
					return false
				} else {
					expectation = expectation[:len(expectation)-1]
				}
			} else if expectation[len(expectation)-1] == '[' {
				if s[i] != ']' {
					return false
				} else {
					expectation = expectation[:len(expectation)-1]
				}
			} else if expectation[len(expectation)-1] == '{' {
				if s[i] != '}' {
					return false
				} else {
					expectation = expectation[:len(expectation)-1]
				}
			}
		}
	}
	return true
}

func main() {
	fmt.Println(isValid("([])"))
}
