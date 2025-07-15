package main

import "fmt"

func isValid(word string) bool {
	var flagCh, flagDi, flagVo, flagCo = false, true, false, false
	c := 0
	for _, l := range word {
		c++
		if !flagCh && c > 2 {
			flagCh = true
		}
		if (l < '0' || l > '9') && (l < 'A' || l > 'z') {
			flagDi = false
			break
		}
		if !flagCo && (l >= 'A' && l <= 'z') {
			flagCo = isConstant(l)
		}
		if !flagVo && (l >= 'A' && l <= 'z') {
			flagVo = !isConstant(l)
		}
	}
	return flagCh && flagCo && flagDi && flagVo
}

func isConstant(char rune) bool {
	for _, vowel := range "aeiouAEIOU" {
		if char == vowel {
			return false
		}
	}
	return true
}

func main() {
	fmt.Println(isValid("as/dada1212a"))
}
