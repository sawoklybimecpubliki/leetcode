package main

import (
	"strings"
)

func main() {
	s := "(ed(et(oc))el)"
	var runes []rune
	var interval [2]int
	for true {
		interval[1] = strings.Index(s, ")")
		if interval[1] == -1 {
			println(s)
			println(string(runes))
			return
		}
		interval[0] = strings.LastIndex(s[:interval[1]], "(")
		println("interval:", interval[0], interval[1])
		runes = []rune(s[interval[0]+1 : interval[1]])
		for k, j := 0, len(runes)-1; k < j; k, j = k+1, j-1 {
			runes[k], runes[j] = runes[j], runes[k]
		}
		if interval[0] == 0 {
			if interval[1] == len(s)-1 {
				s = string(runes)
			} else {
				s = string(runes) + s[interval[1]+1:]
			}
		} else {
			if interval[1] == len(s)-1 {
				s = s[:interval[0]] + string(runes)
			} else {
				s = s[:interval[0]] + string(runes) + s[interval[1]+1:]
			}
		}
	}
}
