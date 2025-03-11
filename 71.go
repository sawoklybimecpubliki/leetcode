package main

import "strings"

func main() {
	path := "/a/./b/../../c/"
	var stack []string
	for _, word := range strings.Split(path, "/") {
		if word == ".." {
			stack = stack[:len(stack)-1]
		} else if word == "." || word == "" {
			continue
		} else {
			stack = append(stack, word)
		}
	}
	path = ""
	for len(stack) > 0 {
		println("current: ", stack[len(stack)-1])
		path = "/" + stack[len(stack)-1] + path
		stack = stack[:len(stack)-1]
	}
	if len(path) == 0 {
		path = "/"
	}
	println(path)
	/*var prev [2]int
	for i := 0; i < len(path); {
		switch path[i] {
		case '/':
			if i == len(path)-1 && path[i] == '/' {
				if len(path) == 1 {
					println(path)
					return
				}
				path = path[:i]
				break
			}
			if path[i+1] == '/' {
				path = path[:i+1] + path[i+2:]
			} else {
				prev[0], prev[1] = prev[1], i
				i++
			}
		case '.':
			if i != len(path)-2 && path[i+2] == '.' {
				if path[i+1] == '.' {
					for path[i] == '.' {
						i++
					}
				}
			} else if path[i+1] == '.' {
				if i != len(path)-1 {
					prev[0], prev[1] = strings.LastIndex(path[:prev[1]], "/"), prev[0]
					if prev[0] == -1 {
						prev[0] = 0
					}
					path = path[:prev[0]+1] + path[i+2:]
					i = prev[0]
				} else {
					path = path[:prev[1]+1]
					i = prev[0]
				}
			} else {
				path = path[:prev[1]+1] + path[i+1:]
				i = prev[1]
			}
		default:
			i++
		}
	}
	println(path)*/
}
