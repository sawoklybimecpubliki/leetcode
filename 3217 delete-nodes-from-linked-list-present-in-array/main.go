package main

type ListNode struct {
	Val  int
	Next *ListNode
}

func modifiedList(nums []int, head *ListNode) *ListNode {
	toDelete := make(map[int]struct{}, len(nums))
	for _, v := range nums {
		toDelete[v] = struct{}{}
	}

	dummy := &ListNode{Val: 0, Next: head}
	prev := dummy
	curr := head

	for curr != nil {
		if _, ok := toDelete[curr.Val]; ok {
			prev.Next = curr.Next
		} else {
			prev = curr
		}
		curr = curr.Next
	}
	return dummy.Next
}
