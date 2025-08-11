package main

func minimumRecolors(blocks string, k int) int {
	var m, c, black int
	m = k
	for i := 0; i < len(blocks); i++ {
		for j := i; j < len(blocks); j++ {
			if blocks[j] == 'W' {
				c++
				black++
			} else {
				black++
			}
			if black == k {
				if m > c {
					m = c
				}
				break
			}
		}
		black, c = 0, 0
	}
	return m
}

func main() {
	println(minimumRecolors("BWWWBB", 6))
}
