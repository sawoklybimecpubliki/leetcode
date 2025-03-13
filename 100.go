package main

import "fmt"

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

func isSameTree(p *TreeNode, q *TreeNode) bool {
	if p.Val != q.Val {
		return false
	}
	if p.Left != nil && q.Left != nil {
		if !isSameTree(p.Left, q.Left) {
			return false
		}
	}
	if (p.Left == nil && q.Left != nil) || (p.Left != nil && q.Left == nil) {
		return false
	}
	if p.Right != nil && q.Right != nil {
		if !isSameTree(p.Right, q.Right) {
			return false
		}
	}
	if (p.Left == nil && q.Left != nil) || (p.Left != nil && q.Left == nil) {
		return false
	}
	return true
}

func main() {
	tree1 := NewTreeNode([]int{1, 2, 3})
	tree2 := NewTreeNode([]int{1, 2, 3})
	fmt.Println(isSameTree(tree1, tree2))
}
