package main

func mergeTwoLists(list1 *ListNode, list2 *ListNode) *ListNode {
	head := new(ListNode)
	p1, p2, p3 := list1, list2, head
	for p1 != nil && p2 != nil {
		if p1.Val < p2.Val {
			p3.Next = p1
			p3 = p3.Next
			p1 = p1.Next
		} else {
			p3.Next = p2
			p3 = p3.Next
			p2 = p2.Next
		}
	}
	if p1 != nil {
		p3.Next = p1
	}
	if p2 != nil {
		p3.Next = p2
	}
	return head.Next
}
