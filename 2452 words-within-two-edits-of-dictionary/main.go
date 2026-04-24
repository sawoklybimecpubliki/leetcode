package main

import (
	"fmt"
)

func twoEditWords(queries []string, dictionary []string) []string {
	var ans []string
	for _, query := range queries {
		for _, word := range dictionary {
			l := len(query)
			c := 0
			for j := 0; j < l; j++ {
				if query[j] == word[j] {
					c++
				}
			}
			if c >= l-2 {
				ans = append(ans, query)
				break
			}
		}
	}
	return ans
}

func main() {
	queries := []string{"tsl", "sri", "yyy", "rbc", "dda", "qus", "hyb", "ilu", "ahd"}
	dictionary := []string{"uyj", "bug", "dba", "xbe", "blu", "wuo", "tsf", "tga"}
	fmt.Println(twoEditWords(queries, dictionary))
}
