package main

import "fmt"

var dir = [][]int{{-1, 0}, {1, 0}, {0, -1}, {0, 1}}

func search(board [][]byte, word string, point []int) bool {
	m := len(board)
	n := len(board[0])
	for i := 0; i < 4; i++ {
		if word == "" {
			return true
		}
		if (point[0]+dir[i][0] >= 0 && point[1]+dir[i][1] >= 0) && (point[0]+dir[i][0] < m && point[1]+dir[i][1] < n) {
			if word[0] == board[point[0]+dir[i][0]][point[1]+dir[i][1]] {
				board[point[0]+dir[i][0]][point[1]+dir[i][1]] = ' '
				if search(board, word[1:], []int{point[0] + dir[i][0], point[1] + dir[i][1]}) {
					return true
				}
				board[point[0]+dir[i][0]][point[1]+dir[i][1]] = word[0]
			}
		}
	}
	return false
}

func exist(board [][]byte, word string) bool {
	n := len(board)
	m := len(board[0])
	c := 0
	for i := 0; i < n; i++ {
		for j := 0; j < m; j++ {
			if word[c] == board[i][j] {
				c++
				board[i][j] = ' '
				if search(board, word[1:], []int{i, j}) {
					return true
				} else {
					board[i][j] = word[0]
					c--
				}
				continue
			}
		}
	}
	if c == 0 {
		return false
	}
	return true
}

func main() {
	fmt.Println(exist([][]byte{{'a', 'a', 'b', 'a', 'a', 'b'}, {'a', 'a', 'b', 'b', 'b', 'b', 'a'}, {'a', 'a', 'a',
		'a', 'b', 'a'}, {'b', 'a', 'b', 'b', 'a', 'b'}, {'a', 'b', 'b', 'a', 'b', 'a'}, {'b', 'a', 'a', 'a', 'a', 'b'}}, "bbbaabbbbbab"))
}
