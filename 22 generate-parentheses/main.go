package main

import "fmt"

func generateParenthesis(n int) []string {
	out := []string{}
	stack := make([]byte, 0, 2*n)

	var backtrack func(open, close int)
	backtrack = func(open, close int) {
		if len(stack) == 2*n {
			out = append(out, string(stack))
			return
		}

		if open < n {
			stack = append(stack, '(')
			backtrack(open+1, close)
			stack = stack[:len(stack)-1]
		}

		if close < open {
			stack = append(stack, ')')
			backtrack(open, close+1)
			stack = stack[:len(stack)-1]
		}
	}

	backtrack(0, 0)
	return out
}

func main() {
	fmt.Println(generateParenthesis(3))
}
