package main

import (
	"fmt"
	"strconv"
)

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func NewTreeNode(data []int) *TreeNode {
	return new(TreeNode).insertNode(data, 0)
}

func (r *TreeNode) insertNode(data []int, i int) *TreeNode {
	var root *TreeNode
	if i >= len(data) || data[i] == -1 {
		return root
	}
	root = &TreeNode{Val: data[i]}
	root.Left = root.insertNode(data, 2*i+1)
	root.Right = root.insertNode(data, 2*i+2)
	return root
}

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
	fmt.Println(sumNumbers(tree))
}
