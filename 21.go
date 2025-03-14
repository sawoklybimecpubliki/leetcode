package main

import "fmt"

type ListNode struct {
	Val  int
	Next *ListNode
}

func NewList(num []int) *ListNode {
	var startNode, prevNode, iNode *ListNode
	if len(num) == 0 {
		return &ListNode{}
	}
	startNode = &ListNode{Val: num[0], Next: nil}
	prevNode = startNode
	for _, i := range num[1:] {
		iNode = &ListNode{Val: i, Next: nil}
		prevNode.Next = iNode
		prevNode = iNode
	}
	return startNode
}

func mergeTwoLists(list1 *ListNode, list2 *ListNode) *ListNode {
	pointerList1 := list1
	pointerList2 := list2

	for pointerList2 != nil {

		if pointerList1.Val <= pointerList2.Val && (pointerList1.Next.Val >= pointerList2.Val || pointerList1.Next == nil) {
			tmp := &ListNode{pointerList2.Val, nil}

			if pointerList1.Next != nil {
				tmp.Next = pointerList1.Next
			}

			pointerList1.Next = tmp
			pointerList2 = pointerList2.Next
		} else {
			pointerList1 = pointerList1.Next
		}
	}

	return list1
}

func main() {
	list1 := NewList([]int{1, 2, 4})
	list2 := NewList([]int{1, 3, 4})

	for i := mergeTwoLists(list1, list2); i != nil; i = i.Next {
		fmt.Println(i.Val)
	}
}

/*pointer := list1

for i := 1; i < len(arr1); i++ {
	pointer.Next = &ListNode{Val: arr1[i], Next: nil}
	pointer = pointer.Next
}*/

/*solution 1
func mergeTwoLists(list1 *ListNode, list2 *ListNode) *ListNode {
    var outList = &ListNode{}
	pointer := outList
	for list1 != nil || list2 != nil {
		if list1 != nil && (list2 == nil || list1.Val < list2.Val) {
			node := &ListNode{list1.Val, nil}
			pointer.Next = node
			pointer = pointer.Next
			list1 = list1.Next
		} else {
			node := &ListNode{list2.Val, nil}
			pointer.Next = node
			pointer = pointer.Next
			list2 = list2.Next
		}
	}
	return outList.Next
}
*/
