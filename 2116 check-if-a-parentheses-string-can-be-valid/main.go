package main

func main() {
	s := "(()())"
	locked := "111111"
	n := len(s)
	if n%2 != 0 {
		println(false)
	}

	openIndices := []int{}
	unlockedIndices := []int{}

	for i := 0; i < n; i++ {
		if locked[i] == '0' {
			unlockedIndices = append(unlockedIndices, i)
		} else if s[i] == '(' {
			openIndices = append(openIndices, i)
		} else if s[i] == ')' {
			if len(openIndices) > 0 {
				openIndices = openIndices[:len(openIndices)-1]
			} else if len(unlockedIndices) > 0 {
				unlockedIndices = unlockedIndices[:len(unlockedIndices)-1]
			} else {
				println(false)
			}
		}
	}

	for len(openIndices) > 0 && len(unlockedIndices) > 0 && openIndices[len(openIndices)-1] < unlockedIndices[len(unlockedIndices)-1] {
		openIndices = openIndices[:len(openIndices)-1]
		unlockedIndices = unlockedIndices[:len(unlockedIndices)-1]
	}

	println(len(openIndices) == 0 && len(unlockedIndices)%2 == 0)
}
