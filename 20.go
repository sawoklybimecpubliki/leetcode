package main

func main() {
	s := "([])"
	var expectation []uint8
	for i := 0; i < len(s); i++ {
		if s[i] == '(' || s[i] == '[' || s[i] == '{' {
			expectation = append(expectation, s[i])
		}
		if s[i] == ')' || s[i] == ']' || s[i] == '}' {
			if expectation[len(expectation)-1] == '(' {
				if s[i] != ')' {
					println("false")
				} else {
					expectation = expectation[:len(expectation)-1]
				}
			} else if expectation[len(expectation)-1] == '[' {
				if s[i] != ']' {
					println("false")
				} else {
					expectation = expectation[:len(expectation)-1]
				}
			} else if expectation[len(expectation)-1] == '{' {
				if s[i] != '}' {
					println("false")
				} else {
					expectation = expectation[:len(expectation)-1]
				}
			}
		}
	}
	/*
		if brackets['('] == brackets[')'] {
			if brackets['['] == brackets[']'] {
				if brackets['{'] == brackets['}'] {
					println("true")
					return
				}
			}
		}
		println("false")*/
}
