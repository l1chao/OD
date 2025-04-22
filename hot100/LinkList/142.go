package main

func detectCycle(head *ListNode) *ListNode {
	slow, quick := head, head
	var meet *ListNode
	for quick != nil && quick.Next != nil {
		quick = quick.Next.Next
		slow = slow.Next

		if quick == slow {
			meet = quick
			break
		}
	}
	if meet == nil { //如果没有环
		return nil
	}

	p := head
	for {
		if p == meet {
			return p
		}
		p = p.Next
		meet = meet.Next
	}
}
