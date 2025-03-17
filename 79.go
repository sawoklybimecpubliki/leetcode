package main

import "fmt"

func search(board [][]byte, word string, point []int) bool {
	m := len(board)
	n := len(board[0])
	var dir = [][]int{{-1, 0}, {1, 0}, {0, -1}, {0, 1}}

	if word == "" {
		return true
	}

	if point[0] < 0 || point[1] < 0 || point[0] >= m || point[1] >= n || word[0] != board[point[0]][point[1]] {
		return false
	}

	board[point[0]][point[1]] = ' '
	for i := 0; i < 4; i++ {
		dx, dy := point[0]+dir[i][0], point[1]+dir[i][1]

		if search(board, word[1:], []int{dx, dy}) {
			return true
		}

	}
	board[point[0]][point[1]] = word[0]
	return false
}

func exist(board [][]byte, word string) bool {
	n := len(board)
	m := len(board[0])
	for i := 0; i < n; i++ {
		for j := 0; j < m; j++ {
			if search(board, word, []int{i, j}) {
				return true
			}
		}
	}
	return false
}

func main() {
	fmt.Println(exist([][]byte{{'a'}}, "b"))
	/*fmt.Println(exist([][]byte{
	{'A', 'B', 'C', 'E'},
	{'S', 'F', 'C', 'S'},
	{'A', 'D', 'E', 'E'}}, "ABCCED"))*/
	/*fmt.Println(exist([][]byte{{'a', 'a', 'b', 'a', 'a', 'b'}, {'a', 'a', 'b', 'b', 'b', 'b', 'a'}, {'a', 'a', 'a',
	'a', 'b', 'a'}, {'b', 'a', 'b', 'b', 'a', 'b'}, {'a', 'b', 'b', 'a', 'b', 'a'}, {'b', 'a', 'a', 'a', 'a', 'b'}}, "bbbaabbbbbab"))*/
}
