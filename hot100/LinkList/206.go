package main

func reverseList(head *ListNode) *ListNode {
	if head == nil {
		return nil
	}

	blank := new(ListNode)
	p := head
	for p != nil {
		tmp := p
		p = p.Next
		tmp.Next = blank.Next
		blank.Next = tmp
	}
	return blank.Next
}
