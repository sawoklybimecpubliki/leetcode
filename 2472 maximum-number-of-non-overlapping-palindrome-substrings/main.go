package main

import "fmt"

func maxPalindromes(s string, k int) (count int) {
	n := len(s)
	if k == 1 {
		return n
	}
	for i := 0; i+k <= n; i++ {
		if isPalindrome(s, i, i+k-1) {
			count++
			i += k - 1
		} else if i+k < n && isPalindrome(s, i, i+k) {
			count++
			i += k
		}
	}
	return count
}

func isPalindrome(s string, l, r int) bool {
	for l < r && s[l] == s[r] {
		l, r = l+1, r-1
	}
	return l >= r
}

func main() {

	fmt.Println(maxPalindromes("fttfjofpnpfydwdwdnns", 2))
}
