package main

import "fmt"

func clearStars(s string) string {
	a := 'a'
	freq := make([][]int, 26)
	for i := range freq {
		freq[i] = []int{}
	}

	for i, ch := range s {
		if ch == '*' {
			for j := 0; j < 26; j++ {
				if len(freq[j]) > 0 {
					freq[j] = freq[j][:len(freq[j])-1]
					break
				}
			}
		} else {
			freq[ch-a] = append(freq[ch-a], i)
		}
	}

	keep := make([]bool, len(s))
	for _, indexes := range freq {
		for _, id := range indexes {
			keep[id] = true
		}
	}

	result := []rune{}
	for i, ch := range s {
		if keep[i] {
			result = append(result, ch)
		}
	}

	return string(result)
}

func main() {
	fmt.Println(clearStars("aaba*"))
}
