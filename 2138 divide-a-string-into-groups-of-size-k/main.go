package main

import "fmt"

func divideString(s string, k int, fill byte) []string {
	var out []string

	if x := len(s) % k; x == 0 {
		for i := 0; i < len(s); i += k {
			out = append(out, s[i:i+k])
		}
	} else {
		for i := 0; i < len(s)-x; i += k {
			out = append(out, s[i:i+k])
		}
		str := s[len(s)-x:]
		for i := 0; i < k-x; i++ {
			str += string(fill)
		}
		out = append(out, str)
	}
	return out
}

func main() {
	fmt.Println(divideString("ctoyjrwtngqwt", 8, 'n'))
}
