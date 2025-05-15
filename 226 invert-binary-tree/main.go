package main

import "fmt"

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

var DefaultValue int = -1024

func InsertNodeToTree(tree *TreeNode, node *TreeNode) {
	if tree == nil {
		return
	}
	if tree.Val == DefaultValue {
		tree.Val = node.Val
		return
	}
	if node.Val > tree.Val {
		if tree.Right == nil {
			tree.Right = &TreeNode{Val: DefaultValue}
		}
		InsertNodeToTree(tree.Right, node)
	}
	if node.Val < tree.Val {
		if tree.Left == nil {
			tree.Left = &TreeNode{Val: DefaultValue}
		}
		InsertNodeToTree(tree.Left, node)
	}
}

func InitTree(values ...int) (root *TreeNode) {
	rootNode := TreeNode{Val: DefaultValue, Right: nil, Left: nil}
	for _, value := range values {
		node := TreeNode{Val: value}
		InsertNodeToTree(&rootNode, &node)
	}
	return &rootNode
}

func invert(tree *TreeNode) {
	if tree == nil {
		return
	}

	if tree.Right != nil {
		invert(tree.Right)
	}

	if tree.Left != nil {
		invert(tree.Left)
	}

	tree.Right, tree.Left = tree.Left, tree.Right
}

func invertTree(root *TreeNode) *TreeNode {
	printTree(root)
	fmt.Println("---------------------------")
	invert(root)
	printTree(root)
	return root
}

func printTree(root *TreeNode) {
	if nil == root {
		return
	}
	printTree(root.Left)
	fmt.Printf("%d ", root.Val)
	printTree(root.Right)
}

func main() {
	invertTree(InitTree(4, 2, 7, 1, 3, 6, 9))
}
