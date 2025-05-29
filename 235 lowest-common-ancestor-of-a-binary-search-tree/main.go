package main

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

	for i < len(values) {
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

func lowestCommonAncestor(root, p, q *TreeNode) *TreeNode {
	for root != nil {
		if p.Val < root.Val && q.Val < root.Val {
			root = root.Left
		} else if p.Val > root.Val && q.Val > root.Val {
			root = root.Right
		} else {
			return root
		}
	}
	return nil
}
