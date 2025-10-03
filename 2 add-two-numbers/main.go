package main

import "fmt"

type ListNode struct {
	Val  int
	Next *ListNode
}

func constructor(nums []int) *ListNode {
	head := &ListNode{nums[0], nil}
	cur := head
	for i := 1; i < len(nums); i++ {
		node := &ListNode{Val: nums[i]}
		cur.Next = node
		cur = cur.Next
	}
	return head
}

func addTwoNumbers(l1 *ListNode, l2 *ListNode) *ListNode {
	dummy := &ListNode{}
	temp := dummy
	carry := 0

	for l1 != nil || l2 != nil || carry != 0 {
		val1 := 0
		if l1 != nil {
			val1 = l1.Val
			l1 = l1.Next
		}

		val2 := 0
		if l2 != nil {
			val2 = l2.Val
			l2 = l2.Next
		}

		sum := val1 + val2 + carry
		carry = sum / 10
		temp.Next = &ListNode{Val: sum % 10}
		temp = temp.Next
	}

	return dummy.Next
}

func main() {
	var l1, l2 *ListNode
	l1 = constructor([]int{1, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 1})
	fmt.Println(l1)
	l2 = constructor([]int{5, 6, 4})
	fmt.Println(l2)
	res := addTwoNumbers(l1, l2)
	fmt.Println(res)
}
