package main

import (
	"fmt"
	"strconv"
)

func recursive(root *TreeNode, path string, result *int) {
	if root == nil {
		return
	}
	path += strconv.Itoa(root.Val)
	if root.Left != nil {
		recursive(root.Left, path, result)
	}
	if root.Right != nil {
		recursive(root.Right, path, result)
	}
	if root.Left == nil && root.Right == nil {
		tmp, _ := strconv.Atoi(path)
		*result += tmp
	}
}

func sumNumbers(root *TreeNode) int {
	var result int
	recursive(root, "", &result)
	return result
}

func main() {
	tree := NewTreeNode([]int{4, 9, 0, 5, 1})
	fmt.Println(tree)
	//fmt.Println(sumNumbers(tree))
}
