package main

import "fmt"

func pushDominoes(dominoes string) string {
	s := []byte(dominoes)
	l := len(s)
	if l == 1 {
		return string(s)
	}
	for i := 0; i < l; i++ {
		switch string(s[i]) {
		case ".":
			k := i
			for k < l-1 && s[k+1] == '.' {
				k++
			}
			if i == 0 && k == l-1 {
				return string(s)
			}
			if i == 0 {
				if s[i+1] == 'L' {
					s[i] = 'L'
				}
				break
			}

			if i == l-1 {
				if s[i-1] == 'R' {
					s[i] = 'R'
				}
				break
			}

			if i == k && s[i-1] == 'R' && s[i+1] == 'L' {
				break
			} else if i != 0 && s[i-1] == 'R' {
				s[i] = 'R'
			} else if i != l-1 && s[i+1] == 'L' {
				s[i] = 'L'
				i--
			}
			if i != k && k != l-1 && s[k+1] == 'L' {
				s[k] = 'L'
				i--
			}

		}
	}

	return string(s)
}

func main() {
	fmt.Println(pushDominoes("RR.L"))

}
