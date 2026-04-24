package main

import "fmt"

func furthestDistanceFromOrigin(moves string) int {
	var out int
	c := 0
	for _, move := range moves {
		if move == 'L' {
			c--
		}
		if move == 'R' {
			c++
		}
		if move == '_' {
			out++
		}
	}
	if c < 0 {
		c = c * -1
		return out + c
	}
	return out + c
}

/*
	m := make(map[int32]int)
	for _, move := range moves {
		m[move]++
	}
	if m['L'] > m['R'] {
		out = m['L'] - m['R'] + m['_']
		return out
	}
	out = m['R'] - m['L'] + m['_']
*/

func main() {
	moves := "_______"
	fmt.Println(furthestDistanceFromOrigin(moves))
}
