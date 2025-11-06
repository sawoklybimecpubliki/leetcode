package main

import "fmt"

func hasSameDigits(s string) bool {
	b := []byte(s)
	length := len(b)
	for length > 2 {
		for i := 0; i < length-1; i++ {
			b[i] = ((b[i]-'0')+(b[i+1]-'0'))%10 + '0'
		}
		length--
	}
	return b[0] == b[1]
}

func main() {
	fmt.Println(hasSameDigits("3902"))
}
