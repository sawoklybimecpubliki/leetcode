package main

import "fmt"

type ListNode struct {
	Val  int
	Next *ListNode
}

func (n *ListNode) Add(val int) {
	n.Val = val
	n.Next = &ListNode{}
}

func getDecimalValue(head *ListNode) int {
	var num []int
	var res int
	for node := head; node.Next != nil; node = node.Next {
		num = append(num, node.Val)
	}
	j := 0
	fmt.Println(num)
	for i := len(num) - 1; i >= 0; i-- {
		res += num[j] * pow(i)
		j++
	}
	return res
}

func pow(i int) int {
	num := 1
	for range i {
		num *= 2
	}
	return num
}

func main() {
	head := &ListNode{}
	input := []int{1, 0, 1, 1, 1}
	tmp := head
	for _, i := range input {
		tmp.Add(i)
		tmp = tmp.Next
	}
	fmt.Println(getDecimalValue(head))
}
