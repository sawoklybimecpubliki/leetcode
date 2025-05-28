package main

import "fmt"

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func BuildBinaryTree(values []interface{}) *TreeNode {
	if len(values) == 0 || values[0] == nil {
		return nil
	}

	var queue []*TreeNode
	head := new(TreeNode)
	queue = append(queue, head)
	i := 0

	for i < len(values) && values[i] != nil {
		val := values[i].(int)
		current := queue[0]
		queue = queue[1:]

		current.Val = val
		leftIndex := i*2 + 1
		rightIndex := leftIndex + 1

		if leftIndex < len(values) && values[leftIndex] != nil {
			newNode := new(TreeNode)
			current.Left = newNode
			queue = append(queue, current.Left)
		}

		if rightIndex < len(values) && values[rightIndex] != nil {
			newNode := new(TreeNode)
			current.Right = newNode
			queue = append(queue, current.Right)
		}

		i++
	}

	return head
}

func isBalanced(root *TreeNode) bool {
	if root == nil {
		return true
	}

	left := maxDepth(root.Left)
	right := maxDepth(root.Right)

	if abs(left-right) > 1 {
		return false
	}
	return isBalanced(root.Left) && isBalanced(root.Right)
}

func abs(a int) int {
	if a > 0 {
		return a
	}
	return -a
}

func maxDepth(root *TreeNode) int {
	if root == nil {
		return 0
	}

	left := maxDepth(root.Left)
	right := maxDepth(root.Right)

	return max(left, right) + 1
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func main() {
	root := BuildBinaryTree([]interface{}{1, 2, 2, 3, 3, nil, nil, 4, 4})
	fmt.Println(isBalanced(root))
}
