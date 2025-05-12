package main

import (
	"fmt"
)

func findEvenNumbers(digits []int) []int {
	var cnts [10]int
	for _, d := range digits {
		cnts[d]++
	}

	ans := make([]int, 0)
	for i := 1; i <= 9; i++ {
		if cnts[i] == 0 {
			continue
		}
		cnts[i]--
		for j := 0; j <= 9; j++ {
			if cnts[j] == 0 {
				continue
			}
			cnts[j]--
			for k := 0; k <= 9; k += 2 {
				if cnts[k] == 0 {
					continue
				}
				ans = append(ans, i*100+j*10+k)
			}
			cnts[j]++
		}
		cnts[i]++
	}
	return ans
}

/*func findEvenNumbers(digits []int) []int {
	var result, unique []int

	permute(digits, []int{}, &result)
	slices.Sort(result)
	for _, n := range result {
		f := true
		for _, u := range unique {
			if u == n {
				f = false
				break
			}
		}
		if f {
			unique = append(unique, n)
		}
		f = true
	}
	return unique
}

func permute(nums []int, out []int, res *[]int) {
	if len(out) == 3 {
		numb := out[0]*100 + out[1]*10 + out[2]
		if numb%2 == 0 {
			*res = append(*res, numb)
		}
		return
	}

	for i, n := range nums {

		if len(out) == 0 && n == 0 {
			continue
		}
		out = append(out, n)
		var temp []int
		temp = append(temp, nums[:i]...)
		temp = append(temp, nums[i+1:]...)
		permute(temp, out, res)
		out = out[:len(out)-1]
	}
	return
}*/

func main() {
	fmt.Println(findEvenNumbers([]int{
		2, 2, 8, 8, 2}))

}
