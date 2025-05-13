package main

import "fmt"

func lengthAfterTransformations(s string, t int) int {

	/*	for _, letter := range s {
			if letter == 'z' {
				tmp = append(tmp, 'a', 'b')
			} else {
				tmp = append(tmp, byte(letter+1))
			}
		}

		for i := 0; i < t-1; i++ {
			l := len(tmp)
			for _, letter := range tmp {
				if letter == 'z' {
					tmp = append(tmp, 'a', 'b')
				} else {
					tmp = append(tmp, letter+1)
				}
			}
			tmp = tmp[l:]
			fmt.Println(string(tmp))
		}*/

	/*var time []int
	result := len(s)
	for _, letter := range s {
		time = append(time, int('z'-letter))
	}
	for i := 0; i < len(time); i++ {
		if time[i] != 0 {
			tmp := (t - time[i]) / time[i]
			result += tmp
			if tmp > 0 {
				time = append(time, 26-time[i]+'z'-'b')
			} else {
				continue
			}
		} else {
			tmp := t/26 + 1
			result += tmp
			if tmp > 0 {
				time = append(time, 26-time[i]+'z'-'b')
			} else {
				continue
			}
			fmt.Println(t / 26)
		}
	}
	fmt.Println(time)*/

	MOD := 1000000007
	cnt := make([]int, 26)

	for _, char := range s {
		cnt[char-'a']++
	}

	for j := 0; j < t; j++ {
		tmp := make([]int, 26)
		for i := 0; i < 26; i++ {
			if i == 25 {
				tmp[0] = (tmp[0] + cnt[i]) % MOD
				tmp[1] = (tmp[1] + cnt[i]) % MOD
			} else {
				tmp[i+1] = (tmp[i+1] + cnt[i]) % MOD
			}
		}
		cnt = tmp
	}

	len := 0
	for _, c := range cnt {
		len = (len + c) % MOD
	}

	return len
}

func main() {
	fmt.Println(lengthAfterTransformations(
		"jqktcurgdvlibczdsvnsg", 7517))
}
