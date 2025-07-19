package main

import (
	"fmt"
	"sort"
	"strings"
)

func removeSubfolders(folder []string) []string {
	sort.Strings(folder)
	ans := []string{folder[0]}
	for i := 1; i < len(folder); i++ {
		lastFolder := ans[len(ans)-1] + "/"

		if !strings.HasPrefix(folder[i], lastFolder) {
			ans = append(ans, folder[i])
		}
	}

	return ans
}

func main() {
	fmt.Println(removeSubfolders([]string{"/a", "/a/b", "/c/d", "/c/d/e", "/c/f"}))
}
