package main

import "fmt"

func generate(numRows int) [][]int {
	var out [][]int
	var tmp []int
	for i := 0; i < numRows; i++ {
		tmp = append(tmp, 1)
		if i > 1 {
			for j := 1; j < len(tmp)-1; j++ {
				if i == len(tmp)-2 {
					tmp[j] = out[i][j] + out[i-1][j+1]
				} else {
					tmp[j] = out[i-1][j-1] + out[i-1][j]
				}
			}
		}
		merge := make([]int, len(tmp))
		copy(merge, tmp)
		out = append(out, merge)
	}
	return out
}

func main() {
	fmt.Println(generate(5))
}
