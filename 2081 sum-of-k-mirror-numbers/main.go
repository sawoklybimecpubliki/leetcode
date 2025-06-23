package main

import (
	"fmt"
	"strconv"
)

func kMirror(k int, n int) int64 {
	var sum int64
	cur := make([]int64, n)
	cur[0] = 1
	for i := 0; i < n; {
		numStr := strconv.Itoa(int(cur[i]))
		if isPalindrome(numStr) {

			str := strconv.FormatInt(cur[i], k)
			if isPalindrome(str) {
				fmt.Println("cur:", cur[i], str)
				if i < n-1 {
					cur[i+1] = cur[i] + 1
				}
				i++
			} else {
				cur[i]++
			}
		} else {
			cur[i]++
		}
	}
	for _, num := range cur {
		sum += num
	}
	return sum
}

func isPalindrome(s string) bool {
	if s == " " {
		return true
	}
	//s = strings.ToLower(s)
	for i, j := 0, len(s)-1; i <= j; {
		if s[i] < '0' || s[i] > '9' {
			i++
			continue
		}
		if s[j] < '0' || s[j] > '9' {
			j--
			continue
		}
		if s[i] != s[j] {
			return false
		}
		i, j = i+1, j-1
	}
	return true
}

func main() {
	fmt.Println(kMirror(5, 20))
}

/*
func solve(pre int64, isOdd bool) int64 {
	tmp := pre
	if isOdd {
		tmp /= 10
	}

	for tmp > 0 {
		pre = pre*10 + (tmp % 10)
		tmp /= 10
	}

	return pre
}

func kthPalindrome(length int) []int64 {
	var ans []int64

	half := (length + 1)/2 - 1
	mn, mx := int64(1), int64(9)
	for i := 0; i < half; i++ {
		mn *= 10
		mx = mx*10 + 9
	}

	val := int64(0)
	for {
		num := mn + val
		if num > mx {
			break
		}
		complete := solve(num, length%2 == 1)
		ans = append(ans, complete)
		val++
	}

	return ans
}

func isPalin(num int64, base int) bool {
	s := ""
	for num > 0 {
		s += string(rune((num % int64(base)) + '0'))
		num /= int64(base)
	}

	i, j := 0, len(s)-1
	for i < j {
		if s[i] != s[j] {
			return false
		}
		i++
		j--
	}

	return true
}

func kMirror(k int, n int) int64 {
	var sum int64 = 0
	count := 0

	for length := 1; length < 12; length++ {
		ele := kthPalindrome(length)

		for _, val := range ele {
			if isPalin(val, k) {
				sum += val
				count++
				if count == n {
					return sum
				}
			}
		}
	}

	return sum
}
*/
