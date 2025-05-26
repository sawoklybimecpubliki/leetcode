package hasCycle

type ListNode struct {
	Val  int
	Next *ListNode
}

func hasCycle(head *ListNode) bool {
	var arr []*ListNode
	if head == nil {
		return false
	}
	for node := head; node.Next != nil; node = node.Next {
		for _, num := range arr {
			if num == node.Next {
				return true
			}
		}
		arr = append(arr, node)
	}
	return false
}

/*
func hasCycle(head *ListNode) bool {
    nodes := make(map[*ListNode]bool)
    for node := head; node != nil; node = node.Next{
        if nodes[node] == true{
            return true
        }
        nodes[node] = true
    }
    return false
}
*/
