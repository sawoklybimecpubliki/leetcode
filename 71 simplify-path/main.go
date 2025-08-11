package main

import (
	"fmt"
	"strings"
)

func simplifyPath(path string) string {
	var stack []string
	for _, word := range strings.Split(path, "/") {
		if word == ".." {
			if len(stack) != 0 {
				stack = stack[:len(stack)-1]
			} else {
				continue
			}
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
	return path
}

func main() {
	fmt.Println(simplifyPath("/home/user/Documents/../Pictures"))
}
