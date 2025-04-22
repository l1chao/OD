package main

func swapPairs(head *ListNode) *ListNode {
	z := new(ListNode)
	z.Next = head
	p := z

	for p.Next != nil && p.Next.Next != nil {
		tmp1 := p.Next
		tmp2 := p.Next.Next

		tmp1.Next = tmp2.Next
		tmp2.Next = tmp1
		p.Next = tmp2

		p = p.Next.Next
	}
	return z.Next
}
